package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	t.Run("InvalidKeyError", func(t *testing.T) {
		err := NewInvalidKeyError()
		assert.Equal(t, "invalid key provided", err.Error())
	})

}
