package input

import (
	"testing"

	"github.com/eduardothsantos/go-blog/src/domain/tests"
)

func TestValidateNameField(t *testing.T) {
	t.Run("Valid name", func(t *testing.T) {
		var expectedErr error = nil
		err := ValidateNameField("Test author")
		tests.AssertEquals(t, expectedErr, err)
	})

	t.Run("Invalid name", func(t *testing.T) {
		expectedErr := "name.invalid"
		err := ValidateNameField("12312312")
		tests.AssertEquals(t, expectedErr, err.Error())
	})
}

func TestValidateEmailField(t *testing.T) {
	t.Run("Valid email", func(t *testing.T) {
		var expectedErr error = nil
		err := ValidateEmailField("test@Author.com")
		tests.AssertEquals(t, expectedErr, err)
	})

	t.Run("Invalid name", func(t *testing.T) {
		expectedErr := "email.invalid"
		err := ValidateEmailField("12312312@")
		tests.AssertEquals(t, expectedErr, err.Error())
	})
}
