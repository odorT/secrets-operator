//go:generate mockgen -destination=mocks/mock_services_generated.go -package=mocks . FindingService
package ports

import (
	"secrets-operator/internal/core/domain"
)

type FindingService interface {
	Add(domain.FindingsReport) error
	Notify(finding domain.FindingsReport) error
	GetById(repoId int) (domain.RepoFindings, error)
	GetByName(repoName string) ([]map[string]string, error)
}
