package myalgo

import (
	"github.com/ethereum/go-ethereum/consensus"
	"fmt"
	"context"
)

type API struct {
	chain  consensus.ChainReader
	myAlgo *MyAlgo
}

// PrintBlock retrieves a block and returns its pretty printed form.
func (api *API) EchoNumber(ctx context.Context, number uint64) (uint64, error) {
	fmt.Println("called echo number")
	return number, nil
}