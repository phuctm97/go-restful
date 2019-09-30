package users

import (
	"errors"
)

// Errors.
var (
	ErrUsernameIsTooShort            = errors.New("username must be at least 1-character length")
	ErrUsernameIsTooLong             = errors.New("username must be at most 32-character length")
	ErrUsernameWasUsed               = errors.New("username was used")
	ErrUsernameContainsInvalidChar   = errors.New("username must contain only digits, underscores, dashes, dots and alphabetical letters")
	ErrUsernameBeginsWithInvalidChar = errors.New("username must begin with either underscore or an alphabetical letter")
	ErrEmailIsInvalid                = errors.New("email is invalid")
	ErrEmailIsTooLong                = errors.New("email must be at most 128-character length")
	ErrEmailWasUsed                  = errors.New("email was used")
	ErrFullNameIsTooShort            = errors.New("full name must be at least 1-character length")
	ErrFullNameIsTooLong             = errors.New("full name must be at most 128-character length")
	ErrBioIsTooLong                  = errors.New("bio must be at most 128-character length")
)
