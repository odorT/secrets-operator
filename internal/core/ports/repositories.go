//go:generate mockgen -destination=mocks/mock_repositories_generated.go -package=mocks . FindingsRepository,Notifier
package ports

import (
	"secrets-operator/internal/core/domain"
)

type FindingsRepository interface {
	SaveFindingsReport(findingsReport domain.FindingsReport, collectionName string) error
	GetRepoFindingsById(repoId int, collectionName string) (domain.RepoFindings, error)
	SaveAndUpdateRepoFindingsById(repoFindings domain.RepoFindings, repoId int, collectionName string) error
	GetRepositoriesByName(repoName string, collectionName string) ([]map[string]string, error)
}

type Notifier interface {
	SendMessage(message domain.FindingsReport) error
}
