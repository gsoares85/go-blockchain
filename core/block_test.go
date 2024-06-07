package core

import (
	"github.com/gsoares85/go-blockchain/crypto"
	"github.com/gsoares85/go-blockchain/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSignBlock(t *testing.T) {
	privaKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privaKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privaKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privaKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	b.Sign(privKey)
	assert.Nil(t, b.Sign(privKey))

	return b
}
