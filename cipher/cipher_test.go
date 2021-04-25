package cipher

import (
	"encoding/hex"
	"github.com/herzrasen/memo/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCipher(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		key, err := NewKey()
		assert.NoError(t, err)
		assert.NotEmpty(t, key)
		decoded, err := hex.DecodeString(key.s)
		assert.NoError(t, err)
		assert.Equal(t, key.b, decoded)
	})
	t.Run("load key", func(t *testing.T) {
		key, _ := NewKey()
		loaded, err := LoadKey(key.s)
		assert.NoError(t, err)
		assert.Equal(t, key.b, loaded.b)
	})
	t.Run("invalid key format", func(t *testing.T) {
		_, err := LoadKey("Not a hex key")
		assert.Error(t, err)
	})
	t.Run("invalid key length", func(t *testing.T) {
		key := hex.EncodeToString([]byte("i"))
		_, err := LoadKey(key)
		assert.Equal(t, errors.NewInvalidKeyError(), err)
	})
	t.Run("encrypt/decrypt round trip", func(t *testing.T) {
		input := []byte("Hello World")
		key, _ := NewKey()
		encrypted, err := key.Encrypt(input)
		assert.NoError(t, err)
		plain, err := key.Decrypt(encrypted)
		assert.NoError(t, err)
		assert.Equal(t, input, plain)
	})
	t.Run("illegal key", func(t *testing.T) {
		key := Key{b: []byte("nope")}
		_, err := key.Encrypt([]byte("foo"))
		assert.Error(t, err)
	})
	t.Run("decrypt with illegal key", func(t *testing.T) {
		key := Key{b: []byte("nope")}
		_, err := key.Decrypt([]byte(""))
		assert.Error(t, err)
	})
}
