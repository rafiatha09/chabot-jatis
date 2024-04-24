package err_helper

import "errors"

var (
	ErrRepositoryNotFound = errors.New("repository: the requested data not found")
	ErrRepositoryMongo    = errors.New("repository: something went wrong with mongo db")
)
