package rawpow

import (
	"math/big"
	"time"
)

func (RowPow *RowPow) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	log.Info("will verfiyHeader")
	return nil
}

func (RowPow *RowPow) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	log.Info("will verfiyHeaders")

	abort := make(chan struct{})
	results := make(chan error, len(headers))
	go func() {
		for _, header := range headers {
			err := RowPow.VerifyHeader(chain, header, false)
			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

func (RowPow *RowPow) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	log.Info("will verfiy uncles")
	return nil
}

func (RowPow *RowPow) VerifySeal(chain consensus.ChainReader, header *types.Header) error {

	log.Info("will verfiy VerifySeal")

	return nil

}

func (RowPow *RowPow) Prepare(chain consensus.ChainReader, header *types.Header) error {
	log.Info("will prepare the block")

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = RowPow.CalcDifficulty(chain, header.Time.Uint64(), parent)
	return nil
}

func (RowPow *RowPow) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return calcDifficultyHomestead(time, parent)
}

func (RowPow *RowPow) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {

	log.Info("will Finalize the block")

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	b := types.NewBlock(header, txs, uncles, receipts)
	return b, nil
}

func (RowPow *RowPow) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {

	log.Info("will Seal the block")

	time.Sleep(15 * time.Second)
	header := block.Header()
	header.Nonce, header.MixDigest = getRequiredHeader()
	return block.WithSeal(header), nil
}
