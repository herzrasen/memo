package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/herzrasen/memo/errors"
	"io"
)

const (
	KeyLength int = 32
)

type Key struct {
	s string
	b []byte
}

func NewKey() (*Key, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return &Key{}, err
	}
	key := hex.EncodeToString(bytes)
	return &Key{
		s: key,
		b: bytes,
	}, nil
}

func LoadKey(key string) (*Key, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return &Key{}, err
	}
	if len(bytes) != KeyLength {
		return &Key{}, errors.NewInvalidKeyError()
	}
	return &Key{
		s: key,
		b: bytes,
	}, nil
}

func (k *Key) newAuthenticatedEncryption() (cipher.AEAD, error) {
	block, err := aes.NewCipher(k.b)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func (k *Key) Encrypt(plain []byte) ([]byte, error) {
	encryption, err := k.newAuthenticatedEncryption()
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, encryption.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	ciphertext := encryption.Seal(nonce, nonce, plain, nil)
	return ciphertext, nil
}

func (k *Key) Decrypt(encrypted []byte) ([]byte, error) {
	encryption, err := k.newAuthenticatedEncryption()
	if err != nil {
		return nil, err
	}
	nonceSize := encryption.NonceSize()
	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	return encryption.Open(nil, nonce, ciphertext, nil)
}
