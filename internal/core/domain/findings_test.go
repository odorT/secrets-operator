package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type FindingsReportTestSuite struct {
	suite.Suite
	fReport  FindingsReport
	validate *validator.Validate
}

func TestSuiteFindings(t *testing.T) {
	suite.Run(t, new(FindingsReportTestSuite))
}

func (s *FindingsReportTestSuite) SetupTest() {

	// validator will be used to check struct fields' validation errors.
	// it is created once, and used everywhere.
	s.validate = validator.New()

	// create sample FindingsReport struct for all tests. Because there are validation tags, they also must be tested
	// And to test methods of struct with validations, other fields of FindingsReport should also be valid.
	s.fReport.PipelineID = 1
	s.fReport.RepoName = "test name"
	s.fReport.RepoID = 1
	s.fReport.RepoURL = "https://test.com"
	s.fReport.CommitAuthor = "test author"
	s.fReport.CommitSHA = "a85af84d39a32da2c8eba1d88019079aeb0741b0"
	s.fReport.Timestamp = time.Now()
	s.fReport.Findings = Findings{
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
	}
}

func (s *FindingsReportTestSuite) TestFindingsReport_BuildCommitURLTableDrivenShouldPass() {

	tests := []struct {
		name      string
		repoUrl   string
		commitSha string
		want      string
	}{
		{
			"Valid repoUrl and valid commitSha should pass #1",
			"https://gitlab.com/testing-repo",
			"a85af84d39a32da2c8eba1d88019079aeb0741b0",
			"https://gitlab.com/testing-repo/-/commit/a85af84d39a32da2c8eba1d88019079aeb0741b0",
		},
		{
			"Valid repoUrl and valid commitSha should pass #2",
			"https://gitlab.com/testing-repo/",
			"a85af84d39a32da2c8eba1d88019079aeb0741b0",
			"https://gitlab.com/testing-repo/-/commit/a85af84d39a32da2c8eba1d88019079aeb0741b0",
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// arrange
			s.fReport.RepoURL = tt.repoUrl
			s.fReport.CommitSHA = tt.commitSha

			//act
			err := s.validate.Struct(s.fReport)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					s.T().Fatal("error occurred when validating struct fields", err.Error())
				}
				s.T().Fatal("validation errors", err.(validator.ValidationErrors).Error())
			}

			commitUrl := s.fReport.BuildCommitURL()

			// assert
			s.Equal(tt.want, commitUrl, tt.name)
		})
	}
}

func (s *FindingsReportTestSuite) TestFindingsReport_BuildPipelineURLTableDrivenShouldPass() {

	tests := []struct {
		name       string
		repoUrl    string
		pipelineId int
		want       string
	}{
		{
			"Valid repoUrl and valid commitSha should pass #1",
			"https://gitlab.com/testing-repo",
			1,
			"https://gitlab.com/testing-repo/-/pipelines/1",
		},
		{
			"Valid repoUrl and valid commitSha should pass #2",
			"https://gitlab.com/testing-repo/",
			111111111111,
			"https://gitlab.com/testing-repo/-/pipelines/111111111111",
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// arrange
			s.fReport.RepoURL = tt.repoUrl
			s.fReport.PipelineID = tt.pipelineId

			//act
			err := s.validate.Struct(s.fReport)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					s.T().Fatal("error occurred when validating struct fields", err.Error())
				}
				s.T().Fatal("validation errors", err.(validator.ValidationErrors).Error())
			}

			pipelineUrl := s.fReport.BuildPipelineURL()

			// assert
			s.Equal(tt.want, pipelineUrl, tt.name)
		})
	}
}

func (s *FindingsReportTestSuite) TestFindingsReport_BuildUserURLTableDrivenShouldPass() {

	tests := []struct {
		name         string
		repoUrl      string
		commitAuthor string
		want         string
	}{
		{
			"Valid repoUrl and valid commitSha should pass #1",
			"https://gitlab.com/testing-repo",
			"test-user",
			"https://gitlab.com/testing-repo/test-user",
		},
		{
			"Valid repoUrl and valid commitSha should pass #2",
			"https://gitlab.com/testing-repo/",
			"TEST_USER",
			"https://gitlab.com/testing-repo/TEST_USER",
		},
	}

	for _, tt := range tests {

		s.Run(tt.name, func() {

			// arrange
			s.fReport.RepoURL = tt.repoUrl
			s.fReport.CommitAuthor = tt.commitAuthor

			//act
			err := s.validate.Struct(s.fReport)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					s.T().Fatal("error occurred when validating struct fields", err.Error())
				}
				s.T().Fatal("validation errors", err.(validator.ValidationErrors).Error())
			}

			commitAuthor := s.fReport.BuildUserURL()

			// assert
			s.Equal(tt.want, commitAuthor, tt.name)
		})
	}
}
