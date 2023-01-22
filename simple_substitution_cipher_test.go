package simple_substitution_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomKey(t *testing.T) {
	key := RandomKey()
	assert.Equal(t, 26, len(key))
	t.Log(key)
}

func TestEncrypt(t *testing.T) {
	ciphertext, err := Encrypt("CC11001100", "WDEQRGJIFPMTBKNVHSCAZLUYOX")
	assert.Nil(t, err)
	assert.Equal(t, "EE11001100", ciphertext)
}

func TestDecrypt(t *testing.T) {
	plaintext, err := Decrypt("EE11001100", "WDEQRGJIFPMTBKNVHSCAZLUYOX")
	assert.Nil(t, err)
	assert.Equal(t, "CC11001100", plaintext)
}
