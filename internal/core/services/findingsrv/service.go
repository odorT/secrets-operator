package findingsrv

import (
	"go.uber.org/zap"
	"secrets-operator/internal/core/domain"
	"secrets-operator/internal/core/ports"
	"secrets-operator/internal/errors"
)

type service struct {
	l                  *zap.SugaredLogger
	findingsRepository ports.FindingsRepository
	notifier           ports.Notifier
}

func NewFindingService(l *zap.SugaredLogger, findingsRepository ports.FindingsRepository, notifier ports.Notifier) *service {

	return &service{
		l:                  l,
		findingsRepository: findingsRepository,
		notifier:           notifier,
	}
}

func (srv service) Add(findingsReport domain.FindingsReport) error {

	err := srv.findingsRepository.SaveFindingsReport(findingsReport, "findings")
	if err != nil {
		srv.l.Error(err)
		return errors.ErrCouldNotSaveFindingsReport
	}

	repositoryFindings := domain.RepoFindings{
		Findings: findingsReport.Findings,
		RepoID:   findingsReport.RepoID,
		RepoName: findingsReport.RepoName,
		RepoURL:  findingsReport.RepoURL,
	}

	err = srv.findingsRepository.SaveAndUpdateRepoFindingsById(repositoryFindings, findingsReport.RepoID, "repositories")
	if err != nil {
		srv.l.Error(err)
		return errors.ErrCouldNotSaveAndUpdateRepoFindingsById
	}

	return nil
}

func (srv service) Notify(finding domain.FindingsReport) error {

	err := srv.notifier.SendMessage(finding)
	if err != nil {
		srv.l.Errorln(err)
		return errors.ErrCouldNotNotify
	}

	return nil
}

func (srv service) GetById(repoId int) (domain.RepoFindings, error) {

	repositoryFindings, err := srv.findingsRepository.GetRepoFindingsById(repoId, "repositories")
	if err != nil {
		srv.l.Error(err)
		return domain.RepoFindings{}, errors.ErrCouldNotGetRepoFindingsById
	}

	return repositoryFindings, nil
}

func (srv service) GetByName(repoName string) ([]map[string]string, error) {

	repositories, err := srv.findingsRepository.GetRepositoriesByName(repoName, "repositories")
	if err != nil {
		srv.l.Error(err)
		return nil, errors.ErrCouldNotGetRepositoriesByName
	}

	if len(repositories) == 0 {
		srv.l.Errorln("no repository found with provided name")
		return nil, errors.ErrNoRepositoriesFound
	}

	return repositories, nil
}
