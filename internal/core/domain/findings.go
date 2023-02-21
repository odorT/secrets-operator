package domain

import (
	"fmt"
	_ "github.com/go-playground/validator/v10"
	"time"
)

type Findings []struct {
	Description string    `json:"Description" validate:"required,ascii,max=1000"`
	StartLine   int       `json:"StartLine" validate:"required,number,min=0"`
	EndLine     int       `json:"EndLine" validate:"required,number,min=0"`
	StartColumn int       `json:"StartColumn" validate:"required,number,min=0"`
	EndColumn   int       `json:"EndColumn" validate:"required,number,min=0"`
	Match       string    `json:"Match" validate:"required"`
	Secret      string    `json:"Secret" validate:"required"`
	File        string    `json:"File" validate:"required,ascii,max=200"`
	Commit      string    `json:"Commit" validate:"required,ascii,len=40"`
	Entropy     float64   `json:"Entropy" validate:"required,number,min=0,max=20"`
	Author      string    `json:"Author" validate:"required,ascii,max=200"`
	Email       string    `json:"Email" validate:"required,ascii,email"`
	Date        time.Time `json:"Date" validate:"required"`
	Message     string    `json:"Message" validate:"required,ascii,max=1000"`
	Tags        []string  `json:"Tags" validate:"required"`
	RuleID      string    `json:"RuleID" validate:"required,ascii,max=200"`
	Fingerprint string    `json:"Fingerprint" validate:"required,ascii,max=1000"`
}

type FindingsReport struct {
	PipelineID   int       `json:"pipelineId" validate:"required,number,min=0"`
	RepoName     string    `json:"repoName" validate:"required,ascii,max=1000"`
	RepoID       int       `json:"repoId" validate:"required,number,min=0"`
	RepoURL      string    `json:"repoURL" validate:"required,uri"`
	CommitAuthor string    `json:"commitAuthor" validate:"required,ascii,max=200"`
	CommitSHA    string    `json:"commitSHA" validate:"required,ascii,len=40"`
	Timestamp    time.Time `json:"timestamp" validate:"required"`
	Findings     `json:"findings" validate:"required,dive"`
}

type RepoFindings struct {
	RepoID   int    `json:"repoId" validate:"required,number,min=0"`
	RepoName string `json:"repoName" validate:"required,ascii,max=1000"`
	RepoURL  string `json:"repoURL" validate:"required,uri"`
	Findings `json:"findings" validate:"omitempty,dive"`
}

func (fr *FindingsReport) BuildCommitURL() string {
	return fmt.Sprintf("%s/-/commit/%s", fr.RepoURL, fr.CommitSHA)
}

func (fr *FindingsReport) BuildUserURL() string {
	return fmt.Sprintf("%s/%s", fr.RepoURL, fr.CommitAuthor)
}

func (fr *FindingsReport) BuildPipelineURL() string {
	return fmt.Sprintf("%s/-/pipelines/%d", fr.RepoURL, fr.PipelineID)
}
