package errors

import (
	"errors"
)

var (
	ErrRepositoryNotFound                    = errors.New("repository with given id not found in collection")
	ErrCouldNotSaveFindingsReport            = errors.New("could not save findings report")
	ErrCouldNotSaveAndUpdateRepoFindingsById = errors.New("could not save and update repo findings by id")
	ErrCouldNotNotify                        = errors.New("could not notify")
	ErrCouldNotGetRepoFindingsById           = errors.New("could not get repo findings with provided id")
	ErrCouldNotGetRepositoriesByName         = errors.New("could not get repositories by name")
	ErrNoRepositoriesFound                   = errors.New("no repositories found")
)
