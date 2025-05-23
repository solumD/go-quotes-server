package srverrors

import "errors"

var (
	ErrTextIsEmpty   = errors.New("quote text is empty")
	ErrAuthorIsEmpty = errors.New("quote author is empty")
	ErrInvalidID     = errors.New("invalid quote id")
)
