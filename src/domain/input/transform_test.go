package input

import (
	"testing"

	"github.com/eduardothsantos/go-blog/src/domain/tests"
)

func TestTransformSingleLine(t *testing.T) {
	t.Run("Test transforming single line string", func(t *testing.T) {
		str := "    This is a  test  \n. "
		expectedValue := "This is a test ."
		actualValue := TransformSingleLine(str)

		tests.AssertEquals(t, expectedValue, actualValue)
	})
}
