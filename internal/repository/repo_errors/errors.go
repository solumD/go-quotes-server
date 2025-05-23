package repoerrors

import "errors"

var (
	// ErrAuthorNotExist is returned when quote author does not exist in the database.
	ErrAuthorNotExist = errors.New("author does not exist")
)
