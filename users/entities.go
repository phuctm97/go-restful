package users

import (
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

// ValidateUser validates a user data and returns validation errors or nil.
func ValidateUser(user *User) error {
	errs := make([]error, 0)

	// Validate username.
	if len(user.Username) < 1 {
		errs = append(errs, ErrUsernameIsTooShort)
	}
	if len(user.Username) > 32 {
		errs = append(errs, ErrUsernameIsTooLong)
	}
	if !regexContainsOnlyDigitsUnderscoreDashesDotsAlpha.MatchString(user.Username) {
		errs = append(errs, ErrUsernameContainsInvalidChar)
	}
	if !regexBeginsWithUnderscoreOrAlpha.MatchString(user.Username) {
		errs = append(errs, ErrUsernameBeginsWithInvalidChar)
	}

	// Validate email.
	if !regexIsEmail.MatchString(user.Email) {
		errs = append(errs, ErrEmailIsInvalid)
	}
	if len(user.Email) > 128 {
		errs = append(errs, ErrEmailIsTooLong)
	}

	// Validate full name.
	if len(user.FullName) < 1 {
		errs = append(errs, ErrFullNameIsTooShort)
	}
	if len(user.FullName) > 128 {
		errs = append(errs, ErrFullNameIsTooLong)
	}

	// Validate bio.
	if len(user.Bio) > 256 {
		errs = append(errs, ErrBioIsTooLong)
	}

	// Return errors or nil.
	if len(errs) == 0 {
		return nil
	}
	return common.Errors(errs)
}
