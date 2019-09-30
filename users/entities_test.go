package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		user User
		errs []error
	}{
		{
			User{
				Username: "hasInvalidCh@r",
				Email:    "invalidemail",
				FullName: "Good Full Name",
			},
			[]error{ErrUsernameContainsInvalidChar, ErrEmailIsInvalid},
		},
		{
			User{
				Username: "1beginWithNumber",
				Email:    "good@email.com",
				FullName: "",
			},
			[]error{ErrUsernameBeginsWithInvalidChar, ErrFullNameIsTooShort},
		},
		{
			User{
				Username: "this-username_is.valid.but-too_long",
				Email:    "invalid@email",
				FullName: "Good Full Name",
			},
			[]error{ErrUsernameIsTooLong, ErrEmailIsInvalid},
		},
		{
			User{
				Username: "username.dot-dash_underscore",
				Email:    "good@email.com",
				FullName: "Good Full Name",
				Bio:      "Good Bio",
			},
			nil,
		},
	}

	for _, test := range tests {
		errs := ValidateUser(&test.user)
		if test.errs == nil {
			assert.Nil(t, errs)
		} else {
			assert.ElementsMatch(t, errs, test.errs)
		}
	}
}
