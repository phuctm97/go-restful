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

// ValidateUserUnique validates whether user's username or email is valid and returns errors or nil.
func ValidateUserUnique(userRepo UserRepository, user *User) error {
	errs := make([]error, 0)

	var exists bool
	var err error

	// Validate username.
	exists, err = userRepo.ExistsByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		errs = append(errs, ErrUsernameWasUsed)
	}

	// Validate email.
	exists, err = userRepo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		errs = append(errs, ErrEmailWasUsed)
	}

	// Return errors or nil.
	if len(errs) == 0 {
		return nil
	}
	return common.Errors(errs)
}

// Regular expressions.
var (
	regexContainsOnlyDigitsUnderscoreDashesDotsAlpha = regexp.MustCompile(`^[0-9a-zA-Z_\.\-]*$`)
	regexBeginsWithUnderscoreOrAlpha                 = regexp.MustCompile(`^[a-zA-Z_].*$`)
	regexIsEmail                                     = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)
