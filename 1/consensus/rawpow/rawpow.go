package rawpow

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	errLargeBlockTime    = errors.New("timestamp too big")
	errZeroBlockTime     = errors.New("timestamp equals parent's")
	errTooManyUncles     = errors.New("too many uncles")
	errDuplicateUncle    = errors.New("duplicate uncle")
	errUncleIsAncestor   = errors.New("uncle is ancestor")
	errDanglingUncle     = errors.New("uncle's parent is not ancestor")
	errInvalidDifficulty = errors.New("non-positive difficulty")
	errInvalidMixDigest  = errors.New("invalid mix digest")
	errInvalidPoW        = errors.New("invalid proof-of-work")
)

// New creates a RawPow raw proof-of-work consensus engine
func New(config *params.RawPowConfig, db ethdb.Database) *RawPow {
	// Set any missing consensus parameters to their defaults
	conf := *config
	return &RawPow{
		config: &conf,
		db:     db,
	}
}

// RawPow is the raw proof-of-work consensus engine
type RawPow struct {
	config *params.RawPowConfig // Consensus engine configuration parameters
	db     ethdb.Database       // Database to store and retrieve snapshot checkpoints
	lock   sync.RWMutex         // Protects the signer fields
}

// Author retrieves the Ethereum address of the account that minted the given
// block, which may be different from the header's coinbase if a consensus
// engine is based on signatures.
func (RawPow *RawPow) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyHeader checks whether a header conforms to the consensus rules of a
// given engine. Verifying the seal may be done optionally here, or explicitly
// via the VerifySeal method.
func (RawPow *RawPow) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	// Ensure that the header's extra-data section is of a reasonable size
	// Verify the header's timestamp
	// Verify the block's difficulty based in it's timestamp and parent's difficulty
	// Verify that the gas limit is <= 2^63-1
	// Verify that the gasUsed is <= gasLimit
	// Verify that the gas limit remains within allowed bounds

	// Verify that the block number is parent's +1

	// Verify the engine specific seal securing the block
	if seal {
		if err := RawPow.VerifySeal(chain, header); err != nil {
			return err
		}
	}

	// If all checks passed, validate any special fields for hard forks

	return nil
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications (the order is that of
// the input slice).
func (RawPow *RawPow) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	log.Info("will verfiyHeaders")
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for _, header := range headers {
			err := RawPow.VerifyHeader(chain, header, false)

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (RawPow *RawPow) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	log.Info("will verfiy uncles")
	return nil
}

// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (RawPow *RawPow) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	log.Info("will verfiy VerifySeal")

	// Ensure that we have a valid difficulty for the block
	if header.Difficulty.Sign() <= 0 {
		return errInvalidDifficulty
	}

	// Recompute the digest and PoW value and verify against the header
	return nil
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (RawPow *RawPow) Prepare(chain consensus.ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = RawPow.CalcDifficulty(chain, header.Time.Uint64(), parent)
	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficult
// that a new block should have.
func (RawPow *RawPow) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return calcDifficultyHomestead(time, parent)
}

// Some weird constants to avoid constant memory allocs for them.
var (
	expDiffPeriod = big.NewInt(100000)
	big1          = big.NewInt(1)
	big2          = big.NewInt(2)
	big9          = big.NewInt(9)
	big10         = big.NewInt(10)
	bigMinus99    = big.NewInt(-99)
	big2999999    = big.NewInt(2999999)
)

// calcDifficultyHomestead is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func calcDifficultyHomestead(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        ) + 2^(periodCount - 2)

	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).Set(parent.Time)

	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)

	// 1 - (block_timestamp - parent_timestamp) // 10
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big10)
	x.Sub(big1, x)

	// max(1 - (block_timestamp - parent_timestamp) // 10, -99)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// (parent_diff + parent_diff // 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}
	// for the exponential factor
	periodCount := new(big.Int).Add(parent.Number, big1)
	periodCount.Div(periodCount, expDiffPeriod)

	// the exponential factor, commonly referred to as "the bomb"
	// diff = diff + 2^(periodCount - 2)
	if periodCount.Cmp(big1) > 0 {
		y.Sub(periodCount, big2)
		y.Exp(big2, y, nil)
		x.Add(x, y)
	}
	return x
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (RawPow *RawPow) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	log.Info("will Finalize the block")

	accumulateRewards(chain.Config(), state, header, uncles)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, uncles, receipts), nil
}

// Seal generates a new block for the given input block with the local miner's
// seal place on top.
func (RawPow *RawPow) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	log.Info("will Seal the block")

	time.Sleep(15 * time.Second)

	header := block.Header()
	/*
		runes := []rune(header.ParentHash.String())
		index_in_hash := string(runes[0:3])
		index_in_decimal, _ := strconv.ParseInt(index_in_hash , 0, 64)
		index_in_decimal = index_in_decimal % 10
	*/

	fmt.Print("hash is : ")
	fmt.Print(header.ParentHash.String())

	result := 10
	result_in_float := float64(result)

	header.Nonce, header.MixDigest = getRequiredHeader(result_in_float)
	return block.WithSeal(header), nil
}

func getRequiredHeader(result float64) (types.BlockNonce, common.Hash) {
	return getNonce(result), common.Hash{}
}

func getNonce(result float64) types.BlockNonce {
	var i uint64 = uint64(result)
	var n types.BlockNonce

	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// APIs returns the RPC APIs this consensus engine provides.
func (rawPow *RawPow) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "rawpow",
		Version:   "1.0",
		Service:   &API{chain: chain, rawPow: rawPow},
		Public:    false,
	}}
}

// proof-of-work protocol constants.
var (
	FrontierBlockReward    *big.Int = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
	ByzantiumBlockReward   *big.Int = big.NewInt(3e+18) // Block reward in wei for successfully mining a block upward from Byzantium
	maxUncles                       = 2                 // Maximum number of uncles allowed in a single block
	allowedFutureBlockTime          = 15 * time.Second  // Max time from current time allowed for blocks, before they're considered future blocks
)

// Some weird constants to avoid constant memory allocs for them.
var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

// AccumulateRewards credits the coinbase of the given block with the mining
// reward. The total reward consists of the static block reward and rewards for
// included uncles. The coinbase of each uncle block is also rewarded.
func accumulateRewards(config *params.ChainConfig, state *state.StateDB, header *types.Header, uncles []*types.Header) {
	// Select the correct block reward based on chain progression
	blockReward := FrontierBlockReward
	if config.IsByzantium(header.Number) {
		blockReward = ByzantiumBlockReward
	}
	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		state.AddBalance(uncle.Coinbase, r)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	state.AddBalance(header.Coinbase, reward)
}
