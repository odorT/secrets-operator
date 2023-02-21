package searchHdl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"secrets-operator/config"
	"secrets-operator/internal/core/ports"
)

type httpHandler struct {
	cfg            *config.Config
	l              *zap.SugaredLogger
	validate       *validator.Validate
	findingService ports.FindingService
}

func NewSearchHandler(cfg *config.Config, l *zap.SugaredLogger, findingService ports.FindingService) *httpHandler {

	return &httpHandler{
		cfg:            cfg,
		l:              l,
		validate:       validator.New(),
		findingService: findingService,
	}
}

func (handler *httpHandler) SearchRepositories(c *gin.Context) {

	queryParam := c.Query("query")

	err := handler.validate.Var(queryParam, "required,ascii,max=50")
	if err != nil {
		handler.l.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There is a problem with provided query parameter",
			"error":   err.Error(),
		})
		return
	}

	repoNames, err := handler.findingService.GetByName(queryParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Repositories not found with provided name",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": repoNames,
	})

}
