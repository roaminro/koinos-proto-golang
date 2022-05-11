package canonical

import (
	"bytes"
	"testing"

	"github.com/koinos/koinos-proto-golang/koinos"
	"github.com/koinos/koinos-proto-golang/koinos/protocol"
	"github.com/koinos/koinos-proto-golang/koinos/rpc/block_store"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestCanonicalSerialization(t *testing.T) {
	expectedBytes := []byte{0x0A, 0x03, 0x66, 0x6F, 0x6F, 0x10, 0x01, 0x1a, 0x03, 0x62, 0x61, 0x72}
	bt := koinos.BlockTopology{Id: []byte("foo"), Height: 1, Previous: []byte("bar")}

	// Canonically serialize and test the bytes
	b, err := Marshal(&bt)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(expectedBytes, b))

	// Ensure we can deserialize it
	bt2 := koinos.BlockTopology{}
	err = proto.Unmarshal(b, &bt2)
	assert.Equal(t, err, nil)
	assert.Equal(t, bt.Height, bt2.Height)
	assert.True(t, bytes.Equal(bt.Id, bt2.Id))
	assert.True(t, bytes.Equal(bt.Previous, bt2.Previous))
}

func TestMap(t *testing.T) {
	// Ensure an object with a map gives an error when canonically serializing
	tm := testMap{}
	b, err := Marshal(&tm)
	assert.Nil(t, b)
	assert.ErrorIs(t, err, ErrNoCanonical)

	// Ensure an object with a map nested inside of it gives an error when canonically serializing
	mw := mapWrapper{}
	b, err = Marshal(&mw)
	assert.Nil(t, b)
	assert.ErrorIs(t, err, ErrNoCanonical)
}

func TestUnknownField(t *testing.T) {
	bi := block_store.BlockItem{}
	br := block_store.BlockRecord{}
	br.BlockId = []byte{0x01, 0x02, 0x03}
	br.BlockHeight = 123

	b := protocol.Block{}
	b.Id = []byte{0x04, 0x05, 0x06}

	r := protocol.BlockReceipt{}
	r.Id = []byte{0x04, 0x05, 0x06}

	br.Block = &b
	br.Receipt = &r

	br.PreviousBlockIds = append(br.PreviousBlockIds, []byte{0x07, 0x08, 0x09}, []byte{0x0A, 0x0B, 0x0C})

	brBytes, err := Marshal(&br)
	assert.NoError(t, err)

	err = proto.Unmarshal(brBytes, &bi)
	assert.NoError(t, err)

	biBytes, err := Marshal(&bi)
	assert.NoError(t, err)

	assert.True(t, bytes.Equal(brBytes, biBytes))

	nonCanonBytes := []byte{0x2a, 0x03, 0x07, 0x08, 0x09, 0x2a, 0x03, 0x0a, 0x0b, 0x0c, 0x0a, 0x03, 0x01, 0x02, 0x03, 0x10, 0x7b, 0x1a, 0x05, 0x0a, 0x03, 0x04, 0x05, 0x06, 0x22, 0x05, 0x0a, 0x03, 0x04, 0x05, 0x06}

	bi2 := block_store.BlockItem{}
	err = proto.Unmarshal(nonCanonBytes, &bi2)
	assert.NoError(t, err)

	bi2Bytes, err := Marshal(&bi2)
	assert.NoError(t, err)

	assert.True(t, bytes.Equal(brBytes, bi2Bytes))
}
