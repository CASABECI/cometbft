package factory

import (
	"encoding/hex"
	"strings"

	"github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"
)

func NodeID(str string) types.NodeID {
	id, err := types.NewNodeID(strings.Repeat(str, 2*types.NodeIDByteLength))
	if err != nil {
		panic(err)
	}
	return id
}

func RandomNodeID() types.NodeID {
	id, err := types.NewNodeID(hex.EncodeToString(rand.Bytes( types.NodeIDByteLength)))
	if err != nil {
		panic(err)
	}
	return id
}
