package users

import (
	"errors"
	"regexp"

	"github.com/the-evengers/go-restful/common"
)

// User holds a user's data.
type User struct {
	Username string
	Email    string
	FullName string
	Bio      string
}

// Regular expressions.
var (
	regexContainsOnlyDigitsUnderscoreDashesDotsAlpha = regexp.MustCompile(`^[0-9a-zA-Z_\.\-]*$`)
	regexBeginsWithUnderscoreOrAlpha                 = regexp.MustCompile(`^[a-zA-Z_].*$`)
	regexIsEmail                                     = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

// ValidateUser validates a user data.
func ValidateUser(user *User) error {
	errs := make([]error, 0)

	// Validate username.
	if len(user.Username) < 1 {
		errs = append(errs, errors.New("username must be at least 1-character length"))
	}
	if len(user.Username) > 32 {
		errs = append(errs, errors.New("username must be at most 32-character length"))
	}
	if !regexContainsOnlyDigitsUnderscoreDashesDotsAlpha.MatchString(user.Username) {
		errs = append(errs, errors.New("username must contain only digits, underscores, dashes, dots and alphabetical letters"))
	}
	if !regexBeginsWithUnderscoreOrAlpha.MatchString(user.Username) {
		errs = append(errs, errors.New("username must begin with either underscore or an alphabetical letter"))
	}

	// Validate email.
	if !regexIsEmail.MatchString(user.Email) {
		errs = append(errs, errors.New("email is invalid"))
	}
	if len(user.Email) > 128 {
		errs = append(errs, errors.New("email must be at most 128-character length"))
	}

	// Validate full name.
	if len(user.FullName) < 1 {
		errs = append(errs, errors.New("full name must be at least 1-character length"))
	}
	if len(user.FullName) > 128 {
		errs = append(errs, errors.New("full name must be at most 128-character length"))
	}

	// Validate bio.
	if len(user.Bio) > 256 {
		errs = append(errs, errors.New("bio must be at most 128-character length"))
	}

	// Return errors.
	if len(errs) == 0 {
		return nil
	}
	return common.Errors(errs)
}
