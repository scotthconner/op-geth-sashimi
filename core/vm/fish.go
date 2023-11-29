package vm

type FishStore struct{}

func (c *FishStore) RequiredGas(input []byte) uint64 {
	return 0 // see, cheap gas!
}

func (c *FishStore) Run(context *StatefulPrecompileContext, input []byte) ([]byte, error) {
	return nil, nil
}
