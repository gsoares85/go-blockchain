package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyPairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeyPairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPk := GeneratePrivateKey()
	otherPub := otherPk.PublicKey()

	assert.False(t, sig.Verify(otherPub, msg))
	assert.False(t, sig.Verify(pubKey, []byte("Different msg")))
}
