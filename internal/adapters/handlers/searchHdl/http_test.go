package searchHdl

import (
	"github.com/gin-contrib/cors"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http/httptest"
	"secrets-operator/config"
	"secrets-operator/internal/core/ports/mocks"
	"secrets-operator/internal/errors"
	"testing"
	"time"
)

type SearchHandlerTestSuite struct {
	suite.Suite
	sugaredLogger   *zap.SugaredLogger
	cfg             *config.Config
	ctrl            *gomock.Controller
	setupRouterFunc func() *gin.Engine
}

func TestSuiteSearchHandler(t *testing.T) {
	suite.Run(t, new(SearchHandlerTestSuite))
}

func (s *SearchHandlerTestSuite) SetupTest() {

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

func (s *SearchHandlerTestSuite) TestHttpHandler_SearchRepositories() {

	tests := []struct {
		name                 string
		inputParams          map[string]string
		getByNameReturnValue []map[string]string
		getByNameReturnErr   error
		wantStatusCode       int
		wantResponse         map[string][]map[string]string
	}{
		{
			"valid query parameter and dependencies work should return 200 success",
			map[string]string{"query": "test"},
			[]map[string]string{{"id": "1", "name": "test"}},
			nil,
			200,
			map[string][]map[string]string{"items": {{"id": "1", "name": "test"}}},
		},
		{
			"valid query parameter and dependency returns error should return 404 error",
			map[string]string{"query": "t"},
			nil,
			errors.ErrNoRepositoriesFound,
			404,
			nil,
		},
		{
			"invalid query parameter should return 400 error",
			map[string]string{"q": "test"},
			nil,
			errors.ErrNoRepositoriesFound,
			400,
			nil,
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
				GetByName(gomock.Any()).
				Return(tt.getByNameReturnValue, tt.getByNameReturnErr).
				AnyTimes()

			sut := NewSearchHandler(s.cfg, s.sugaredLogger, mockFindingService)

			// setup new router
			router := s.setupRouterFunc()
			router.GET("/api/v1/search/repos", sut.SearchRepositories)

			// setup request
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/api/v1/search/repos", nil)
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

			// additional assertions
			if tt.wantStatusCode >= 200 && tt.wantStatusCode < 400 {
				resp := map[string][]map[string]string{}

				if err := json.NewDecoder(recorder.Body).Decode(&resp); err != nil {
					s.T().Fatal("could not decode response body into struct.", err)
				}

				assert.Equalf(s.T(), tt.wantResponse, resp, "returned values are not equal. wanted: %s, got: %s", tt.wantResponse, resp)
			}
		})
	}
}
