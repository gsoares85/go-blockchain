package core

import (
	"github.com/gsoares85/go-blockchain/crypto"
	"github.com/gsoares85/go-blockchain/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privaKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privaKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privaKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privaKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
