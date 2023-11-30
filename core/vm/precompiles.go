package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

/*
 * StatefulPrecompileContext
 *
 * This is an object that holds all of the context for execution.
 * It is constructued during VM execution handling (call, execute, etc).
 * This is the main difference between a basic stateless EVM pre-compile
 * and native extensions to the protocol.
 */
type StatefulPrecompileContext struct {
	evm *EVM
}

/*
 * StatefulPrecompileFunction
 *
 * This is a simple interface that enables us to build out many functions and actors
 * and attach them to pre-compile addresses.
 *
 * These objects are created during VM execution and pass in the storage context to
 * a natively compiled contract.
 */
type StatefulPrecompileFunction interface {
	/*
	 * RequiredGas
	 *
	 * Estimates the required gas based on the inputs.
	 *
	 * @param input the call data byte array
	 *
	 * @return the estimated gas price in wei
	 */
	RequiredGas(input []byte) uint64

	/*
	 * Run
	 *
	 * Similar interface to standard geth precompiles, but takes an
	 * additional context object to manage state. Insofar, there is no invariant
	 * protection on this context so it is very much buyer beware.
	 *
	 * @param context the context that was constructed as part of the VM execution
	 * @param input the array of bytes that operate as contract calldata
	 *
	 * @return tuple of the response data and error state.
	 */
	Run(context *StatefulPrecompileContext, input []byte) ([]byte, error)
}

/**
 * StatefulPrecompiledContract
 *
 * A structured mapping of function selectors to precompile functions
 */
type StatefulPrecompileContract struct {
	functions map[common.Hash]StatefulPrecompileFunction
}

/**
 * RunStatefulPrecompiledContract
 *
 * This method is called from the evm execution flow, and combines the
 * execution context to the contract functions. It also estimates and deducts
 * gas costs, or will run out of gas.
 *
 * @param c the execution context, including access to the state
 * @param p the interface for the precompiled contract
 * @param input the byte string that is considered as call data
 * @return a result, the remaining gas of the supplied gas, and any error codes
 */
func RunStatefulPrecompiledContract(c *StatefulPrecompileContext, p *StatefulPrecompileContract, input []byte, suppliedGas uint64) (ret []byte, remainingGas uint64, err error) {
	fmt.Println("what it dooooooooooooo: {}", input[:4])

	// grab the function selector, which is the first four bytes.
	// the contract will hold a map of function interfaces.
	f, ok := p.functions[common.BytesToHash(input[:4])]
	if !ok {
		return nil, suppliedGas, ErrExecutionReverted
	}

	// estimate and cost the gas
	gasCost := f.RequiredGas(input[4:])
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	// execute and return the response and error codes
	output, err := f.Run(c, input[4:])
	return output, suppliedGas, err
}
