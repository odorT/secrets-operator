package findingHdl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"secrets-operator/config"
	"secrets-operator/internal/core/domain"
	"secrets-operator/internal/core/ports"
	"strconv"
	"time"
)

type httpHandler struct {
	cfg            *config.Config
	l              *zap.SugaredLogger
	validate       *validator.Validate
	findingService ports.FindingService
}

func NewFindingsHandler(cfg *config.Config, l *zap.SugaredLogger, findingService ports.FindingService) *httpHandler {

	return &httpHandler{
		cfg:            cfg,
		l:              l,
		validate:       validator.New(),
		findingService: findingService,
	}
}

func (handler *httpHandler) Get(c *gin.Context) {

	repoIdParamString := c.Param("id")

	repoIdParam, err := strconv.Atoi(repoIdParamString)
	if err != nil {
		handler.l.Errorln("could not convert id to int", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not convert parameter from request URI to int",
			"error":   err.Error(),
		})
		return
	}

	err = handler.validate.Var(repoIdParam, "required,number,min=0")
	if err != nil {
		handler.l.Errorln("validation failed for id parameter", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process id parameter in request URI",
			"error":   err.Error(),
		})
		return
	}

	payload, err := handler.findingService.GetById(repoIdParam)
	if err != nil {
		handler.l.Errorln("could not get findings with provided id", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not get findings with provided id",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payload)
}

// Create method
// TODO: find ways to remove type conversions to somewhere else
func (handler *httpHandler) Create(c *gin.Context) {

	payload := &domain.Findings{}

	// decode request body
	err := c.ShouldBindJSON(payload)
	if err != nil {
		handler.l.Errorln("could not decode request body.", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot extract payload from request",
		})
		return
	}

	// checking if the payload is empty slice of findingsReport
	if len(*payload) == 0 {
		handler.l.Errorln("empty request body.", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty findingsReport set",
		})
		return
	}

	// type conversion of parameters from string to int
	repoId, err := strconv.Atoi(c.Query("repoId"))
	if err != nil {
		handler.l.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred when converting repoId parameter to int",
		})
		return
	}

	pipelineId, err := strconv.Atoi(c.Query("pipelineId"))
	if err != nil {
		handler.l.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred when converting pipelineId parameter to int",
		})
		return
	}

	// convert int to int64, because further operations (e.g. converting to Unix time) will require int64
	timestamp, err := strconv.ParseInt(c.Query("timestamp"), 10, 64)
	if err != nil {
		handler.l.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occurred when converting timestamp parameter to int64",
		})
		return
	}

	// we need to create struct matching domain.FindingsReport, where both body and data from query parameters should be set
	findingsReport := domain.FindingsReport{
		PipelineID:   pipelineId,
		RepoName:     c.Query("repoName"),
		RepoID:       repoId,
		RepoURL:      c.Query("repoURL"),
		CommitAuthor: c.Query("commitAuthor"),
		CommitSHA:    c.Query("commitSHA"),
		Timestamp:    time.Unix(timestamp, 0),
		Findings:     *payload,
	}

	err = handler.validate.Struct(findingsReport)
	if err != nil {
		// if there is a problem with validation itself, not user input validation errors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			handler.l.Errorln("Error occurred when validating request parameters", err.Error())
		}

		handler.l.Errorln("Findings validation failed.", err.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed.",
			"error":   err.(validator.ValidationErrors).Error(),
		})
		return
	}

	err = handler.findingService.Add(findingsReport)
	if err != nil {
		handler.l.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save finding, something went wrong",
		})
		return
	}

	if handler.cfg.SlackNotificationEnabled {
		err = handler.findingService.Notify(findingsReport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong when notifying",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
