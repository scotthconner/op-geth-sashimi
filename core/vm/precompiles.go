package vm

/*
 * StatefulPrecompileContext
 *
 * This is an object that holds all of the context for execution.
 * It is constructued during VM execution handling (call, execute, etc).
 * This is the main difference between a basic stateless EVM pre-compile
 * and native extensions to the protocol.
 */
type StatefulPrecompileContext struct {
	blockContext *BlockContext
	stateDB      *StateDB
}

/*
 * StatefulPrecompile
 *
 * This is a simple interface that enables us to build out many functions and actors
 * and attach them to pre-compile addresses.
 *
 * These objects are created during VM execution and pass in the storage context to
 * a natively compiled contract.
 */
type StatefulPrecompile interface {
	/*
	 * RequiredGas
	 *
	 * Estimates the required gas based on the inputs.
	 *
	 * @param input the byte array that could include a function selector as well as raw calldata
	 *              interpretation of these bytes is necessary to take calldata input.
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
