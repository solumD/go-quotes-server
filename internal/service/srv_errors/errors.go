package srverrors

import "errors"

var (
	// ErrTextAndAuthorEmpty is returned when both quote text and author are empty.
	ErrTextAndAuthorEmpty = errors.New("quote text and author are empty")
	// ErrTextIsEmpty is returned when quote text is empty.
	ErrTextIsEmpty = errors.New("quote text is empty")
	// ErrAuthorIsEmpty is returned when quote author is empty.
	ErrAuthorIsEmpty = errors.New("quote author is empty")
	// ErrInvalidID is returned when quote id is less than 0.
	ErrInvalidID = errors.New("invalid quote id")
)
