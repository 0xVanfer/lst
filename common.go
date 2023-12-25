package lst

import (
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Get the call option by block number.
func getCallOpts(block ...int64) *bind.CallOpts {
	// Block length should be 0 or 1. If not 1, use nil as callOpt.
	if len(block) != 1 {
		return nil
	}
	if block[0] == 0 {
		return nil
	}
	return &bind.CallOpts{BlockNumber: types.ToBigInt(block[0])}
}
