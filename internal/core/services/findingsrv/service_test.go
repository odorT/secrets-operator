package findingsrv

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"secrets-operator/internal/core/domain"
	"secrets-operator/internal/core/ports/mocks"
	"secrets-operator/internal/errors"
	"testing"
	"time"
)

type FindingsServiceTestSuite struct {
	suite.Suite
	l    *zap.SugaredLogger
	ctrl *gomock.Controller
}

func TestSuiteFindingService(t *testing.T) {
	suite.Run(t, new(FindingsServiceTestSuite))
}

func (s *FindingsServiceTestSuite) SetupTest() {

	//setup logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	s.l = sugaredLogger

	// setup gomock controller
	s.ctrl = gomock.NewController(s.T())
	defer s.ctrl.Finish()
}

func (s *FindingsServiceTestSuite) TestService_AddTableDriven() {

	tests := []struct {
		name                                     string
		input                                    domain.FindingsReport
		saveFindingsReportReturnValue            error
		saveAndUpdateRepoFindingsByIdReturnValue error
		want                                     error
	}{
		{
			"test #1",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			nil,
			nil,
		},
		{
			"test #2",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			assert.AnError,
			nil,
			errors.ErrCouldNotSaveFindingsReport,
		},
		{
			"test #3",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			assert.AnError,
			errors.ErrCouldNotSaveAndUpdateRepoFindingsById,
		},
		{
			"test #4",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			assert.AnError,
			assert.AnError,
			errors.ErrCouldNotSaveFindingsReport,
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingRepository := mocks.NewMockFindingsRepository(s.ctrl)
			mockNotifier := mocks.NewMockNotifier(s.ctrl)

			mockFindingRepository.
				EXPECT().
				SaveFindingsReport(gomock.Any(), gomock.Any()).
				Return(tt.saveFindingsReportReturnValue).
				AnyTimes()

			mockFindingRepository.
				EXPECT().
				SaveAndUpdateRepoFindingsById(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.saveAndUpdateRepoFindingsByIdReturnValue).
				AnyTimes()

			mockFindingRepository.EXPECT().SaveFindingsReport(tt.input, "test_collection")

			sut := NewFindingService(s.l, mockFindingRepository, mockNotifier)

			// act
			err := sut.Add(tt.input)

			// assert
			assert.Equalf(s.T(), tt.want, err, "assertion failed, wanted: %s, got: %s", tt.want, err)
		})
	}
}

func (s *FindingsServiceTestSuite) TestService_NotifyTableDriven() {

	tests := []struct {
		name                   string
		input                  domain.FindingsReport
		sendMessageReturnValue error
		want                   error
	}{
		{
			"test #1",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			nil,
		},
		{
			"test #2",
			domain.FindingsReport{
				PipelineID:   1,
				RepoName:     "test repo",
				RepoID:       1,
				RepoURL:      "https://test.com",
				CommitAuthor: "test author",
				CommitSHA:    "a85af84d39a32da2c8eba1d88019079aeb0741b0",
				Timestamp:    time.Now(),
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
			assert.AnError,
			errors.ErrCouldNotNotify,
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingRepository := mocks.NewMockFindingsRepository(s.ctrl)
			mockNotifier := mocks.NewMockNotifier(s.ctrl)

			mockNotifier.
				EXPECT().
				SendMessage(gomock.Any()).
				Return(tt.sendMessageReturnValue).
				AnyTimes()

			sut := NewFindingService(s.l, mockFindingRepository, mockNotifier)

			// act
			err := sut.Notify(tt.input)

			// assert
			assert.Equalf(s.T(), tt.want, err, "assertion failed, wanted: %s, got: %s", tt.want, err)
		})
	}
}

func (s *FindingsServiceTestSuite) TestService_GetByIdTableDriven() {

	tests := []struct {
		name                            string
		input                           int
		getRepoFindingsByIdReturnValues domain.RepoFindings
		getRepoFindingsByIdReturnErr    error
		wantValues                      domain.RepoFindings
		wantErr                         error
	}{
		{
			"test #1",
			1,
			domain.RepoFindings{},
			nil,
			domain.RepoFindings{},
			nil,
		},
		{
			"test #2",
			0,
			domain.RepoFindings{},
			assert.AnError,
			domain.RepoFindings{},
			errors.ErrCouldNotGetRepoFindingsById,
		},
		{
			"test #2",
			11111111111111,
			domain.RepoFindings{},
			assert.AnError,
			domain.RepoFindings{},
			errors.ErrCouldNotGetRepoFindingsById,
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingRepository := mocks.NewMockFindingsRepository(s.ctrl)
			mockNotifier := mocks.NewMockNotifier(s.ctrl)

			mockFindingRepository.
				EXPECT().
				GetRepoFindingsById(gomock.Any(), gomock.Any()).
				Return(tt.getRepoFindingsByIdReturnValues, tt.getRepoFindingsByIdReturnErr).
				AnyTimes()

			sut := NewFindingService(s.l, mockFindingRepository, mockNotifier)

			// act
			findings, err := sut.GetById(tt.input)

			// assert
			assert.Equalf(s.T(), tt.wantErr, err, "assertion failed, wanted: %s, got: %s", tt.wantErr, err)
			assert.Equalf(s.T(), tt.wantValues, findings, "assertion failed, wanted: %v, got: %v", tt.wantValues, findings)
		})
	}
}

func (s *FindingsServiceTestSuite) TestService_GetByNameTableDriven() {

	tests := []struct {
		name                              string
		input                             string
		getRepositoriesByNameReturnValues []map[string]string
		getRepositoriesByNameReturnErr    error
		wantValues                        []map[string]string
		wantErr                           error
	}{
		{
			"#1 should pass",
			"test",
			[]map[string]string{{"id": "1", "name": "test"}},
			nil,
			[]map[string]string{{"id": "1", "name": "test"}},
			nil,
		},
		{
			"test #2",
			"test",
			[]map[string]string{},
			assert.AnError,
			nil,
			errors.ErrCouldNotGetRepositoriesByName,
		},
		{
			"test #3",
			"test",
			[]map[string]string{},
			nil,
			[]map[string]string(nil),
			errors.ErrNoRepositoriesFound,
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// because gomock does not support expectation overrides, mock expectations done here, in every loop iteration.
			// when gomock will support expectation overrides, it can be easily moved to SetupTest method of TestSuiteFindingService.

			// arrange
			mockFindingRepository := mocks.NewMockFindingsRepository(s.ctrl)
			mockNotifier := mocks.NewMockNotifier(s.ctrl)

			mockFindingRepository.
				EXPECT().
				GetRepositoriesByName(gomock.Any(), gomock.Any()).
				Return(tt.getRepositoriesByNameReturnValues, tt.getRepositoriesByNameReturnErr).
				AnyTimes()

			sut := NewFindingService(s.l, mockFindingRepository, mockNotifier)

			// act
			findings, err := sut.GetByName(tt.input)

			// assert
			assert.Equalf(s.T(), tt.wantErr, err, "assertion failed, wanted: %s, got: %s", tt.wantErr, err)
			assert.Equalf(s.T(), tt.wantValues, findings, "assertion failed, wanted: %v, got: %v", tt.wantValues, findings)
		})
	}
}
