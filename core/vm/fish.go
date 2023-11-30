package vm

import (
	"github.com/ethereum/go-ethereum/common"
)

var FishStore = StatefulPrecompileContract{
	functions: map[common.Hash]StatefulPrecompileFunction{
		common.HexToHash("0x713e02da"): &CreateSlice{},
	},
}

// //////////////////////////////////////////////////////////////
// createSlice
//
// A user will call this method when they want to create a brand new
// record. The record will come specified, along with the original
// data. This method is considered "payable." Any funds sent
// in with this call will result in the balance being deposited
// into the resulting record's rent account.
//
// @param payload		the raw bytes for the initial contents of the record. could be empty.
//
// @return sliceID 		the resulting slice UUID
// //////////////////////////////////////////////////////////////
type CreateSlice struct{}

func (c *CreateSlice) RequiredGas(input []byte) uint64 {
	return 0 // see, cheap gas!
}

var FishFunctionSelectors = map[string]StatefulPrecompileFunction{}

func (c *CreateSlice) Run(context *StatefulPrecompileContext, input []byte) ([]byte, error) {
	return nil, nil
}
