package core

import (
	"github.com/gsoares85/go-blockchain/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	data := []byte("foo")
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	data := []byte("foo")
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: data,
	}
	signErr := tx.Sign(privKey)
	assert.Nil(t, signErr)
	assert.Equal(t, privKey.PublicKey(), tx.From)
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))

	return tx
}
