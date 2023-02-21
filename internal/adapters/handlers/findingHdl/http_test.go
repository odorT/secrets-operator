package findingHdl

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/cors"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http/httptest"
	"secrets-operator/config"
	"secrets-operator/internal/core/domain"
	"secrets-operator/internal/core/ports/mocks"
	"secrets-operator/internal/errors"
	"testing"
	"time"
)

type FindingsHandlerTestSuite struct {
	suite.Suite
	sugaredLogger   *zap.SugaredLogger
	cfg             *config.Config
	ctrl            *gomock.Controller
	setupRouterFunc func() *gin.Engine
}

func TestSuiteFindingsHandler(t *testing.T) {
	suite.Run(t, new(FindingsHandlerTestSuite))
}

func (s *FindingsHandlerTestSuite) SetupTest() {

	var err error

	//setup sugaredLogger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	s.sugaredLogger = logger.Sugar()

	// setup configs
	s.cfg, err = config.LoadConfig("test")
	if err != nil {
		s.T().Fatalf("cannot load configuration variables. %v", err.Error())
	}
	s.cfg.SlackNotificationEnabled = true

	// setup gomock controller
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()

	// setup router
	s.setupRouterFunc = func() *gin.Engine {
		router := gin.New()
		router.Use(gin.Recovery())
		router.Use(cors.Default())
		router.Use(ginZap.Ginzap(logger, time.RFC3339, true))

		return router
	}
}

func (s *FindingsHandlerTestSuite) TestHttpHandler_Get() {

	tests := []struct {
		name               string
		inputParam         interface{}
		getByIdReturnValue domain.RepoFindings
		getByIdReturnErr   error
		wantStatusCode     int
		wantResponse       domain.RepoFindings
	}{
		{
			"valid path parameter and working dependencies",
			1,
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
			nil,
			200,
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
		},
		{
			"invalid path parameter (negative value)",
			-1,
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
			nil,
			400,
			domain.RepoFindings{},
		},
		{
			"invalid path parameter (nil value)",
			nil,
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
			nil,
			400,
			domain.RepoFindings{},
		},
		{
			"invalid path parameter (string value)",
			"test",
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
			nil,
			400,
			domain.RepoFindings{},
		},
		{
			"empty path parameter",
			"",
			domain.RepoFindings{
				RepoID:   1,
				RepoName: "test",
				RepoURL:  "https://gitlab.com/testing-repo",
				Findings: domain.Findings{
					{
						Description: "test",
						StartLine:   1,
						EndLine:     1,
						StartColumn: 1,
						EndColumn:   1,
						Match:       "test match",
						Secret:      "test secret",
						File:        "test file",
						Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
						Entropy:     3.3822913,
						Author:      "test author",
						Email:       "test@mail.com",
						Date:        time.Now(),
						Message:     "test message",
						Tags:        []string{"test", "tags"},
						RuleID:      "test ruleId",
						Fingerprint: "test fingerprint",
					},
				},
			},
			nil,
			404,
			domain.RepoFindings{},
		},
		{
			"valid path parameter but dependencies return err",
			1,
			domain.RepoFindings{},
			errors.ErrCouldNotGetRepoFindingsById,
			404,
			domain.RepoFindings{},
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingService := mocks.NewMockFindingService(s.ctrl)

			mockFindingService.
				EXPECT().
				GetById(gomock.Any()).
				Return(tt.getByIdReturnValue, tt.getByIdReturnErr).
				AnyTimes()

			sut := NewFindingsHandler(s.cfg, s.sugaredLogger, mockFindingService)

			// setup new router for testing
			router := s.setupRouterFunc()
			router.GET("/api/v1/findings/:id", sut.Get)

			// setup request
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/findings/%v", tt.inputParam), nil)
			request.Header.Set("Content-Type", "application/json")

			// act
			router.ServeHTTP(recorder, request)

			// assert
			assert.Equalf(s.T(), tt.wantStatusCode, recorder.Result().StatusCode, "status codes mismatched. wanted: %d, got: %d", tt.wantStatusCode, recorder.Result().StatusCode)

			// additional assertions
			if tt.wantStatusCode >= 200 && tt.wantStatusCode < 400 {
				resp := domain.RepoFindings{}

				if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
					s.T().Fatal("could not decode response body.", err)
				}

				equal := cmp.Equal(tt.wantResponse, resp)

				assert.Truef(s.T(), equal, "returned values are not equal. wanted: %v, got: %v", tt.getByIdReturnValue, resp)
			}
		})
	}
}

func (s *FindingsHandlerTestSuite) TestHttpHandler_Create() {

	tests := []struct {
		name                string
		inputParams         map[string]string
		inputFindingsReport interface{}
		addReturnErr        error
		notifyReturnErr     error
		wantStatusCode      int
	}{
		{
			"valid request",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			201,
		},
		{
			"valid request but add service return err",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			errors.ErrCouldNotSaveFindingsReport,
			nil,
			500,
		},
		{
			"valid request but add service return err",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			errors.ErrCouldNotSaveAndUpdateRepoFindingsById,
			nil,
			500,
		},
		{
			"valid request but notify service return err",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			errors.ErrCouldNotNotify,
			500,
		},
		{
			"invalid pipelineId query parameter",
			map[string]string{
				"pipelineId":   "-1",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid pipelineId query parameter",
			map[string]string{
				"pipelineId":   "",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid pipelineId query parameter",
			map[string]string{
				"pipelineId":   "test",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid pipelineId query parameter",
			map[string]string{
				"pipelineId":   "§111",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid pipelineId query parameter",
			map[string]string{
				"pipelineId":   "1 1",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoName query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoName query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "§",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoId query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "-44",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoId query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoId query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "test",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoId query parameter",
			map[string]string{
				"pipelineId":   "111",
				"repoName":     "testing repo",
				"repoId":       "§§44",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoId query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1 444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoURL query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https:// gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid repoURL query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid commitAuthor query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid commitAuthor query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "§",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid commitSHA query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8e",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid commitSHA query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid commitSHA query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "§234567891234567891234567891234567891234",
				"timestamp":    "1670071694",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid timestamp query parameter",
			map[string]string{
				"pipelineId":   "1",
				"repoName":     "testing repo",
				"repoId":       "1444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "12345678912345678912345678912345678912341",
				"timestamp":    "",
			},
			domain.Findings{
				{
					Description: "test",
					StartLine:   1,
					EndLine:     1,
					StartColumn: 1,
					EndColumn:   1,
					Match:       "test match",
					Secret:      "test secret",
					File:        "test file",
					Commit:      "a85af84d39a32da2c8eba1d88019079aeb0741b0",
					Entropy:     3.3822913,
					Author:      "test author",
					Email:       "test@mail.com",
					Date:        time.Now(),
					Message:     "test message",
					Tags:        []string{"test", "tags"},
					RuleID:      "test ruleId",
					Fingerprint: "test fingerprint",
				},
			},
			nil,
			nil,
			400,
		},
		{
			"invalid (empty) request body",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{},
			nil,
			nil,
			400,
		},
		{
			"invalid (empty) request body",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			domain.Findings{{}},
			nil,
			nil,
			400,
		},
		{
			"invalid (empty) request body",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			nil,
			nil,
			nil,
			400,
		},
		{
			"invalid (empty) request body",
			map[string]string{
				"pipelineId":   "2",
				"repoName":     "testing repo",
				"repoId":       "444",
				"repoURL":      "https://gitlab.com/testing-repo",
				"commitAuthor": "test user",
				"commitSHA":    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				"timestamp":    "1670071694",
			},
			map[string]string{
				"test": "testing",
			},
			nil,
			nil,
			400,
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingService := mocks.NewMockFindingService(s.ctrl)

			mockFindingService.
				EXPECT().
				Add(gomock.Any()).
				Return(tt.addReturnErr).
				AnyTimes()

			mockFindingService.
				EXPECT().
				Notify(gomock.Any()).
				Return(tt.notifyReturnErr).
				AnyTimes()

			sut := NewFindingsHandler(s.cfg, s.sugaredLogger, mockFindingService)

			// setup new router for testing
			router := s.setupRouterFunc()
			router.POST("/api/v1/findings/upload", sut.Create)

			// prepare request body for post request
			reqBodyBytes := new(bytes.Buffer)
			err := json.NewEncoder(reqBodyBytes).Encode(tt.inputFindingsReport)
			if err != nil {
				s.T().Fatal("could not encode request body for testing.", err)
			}

			// setup request
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/api/v1/findings/upload", reqBodyBytes)
			request.Header.Set("Content-Type", "application/json")

			// add request params
			queryParams := request.URL.Query()
			for k, v := range tt.inputParams {
				queryParams.Add(k, v)
			}

			request.URL.RawQuery = queryParams.Encode()

			// act
			router.ServeHTTP(recorder, request)

			// assert
			assert.Equalf(s.T(), tt.wantStatusCode, recorder.Result().StatusCode, "status codes mismatched. wanted: %d, got: %d", tt.wantStatusCode, recorder.Result().StatusCode)
		})
	}
}
