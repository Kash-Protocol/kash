package randomx

// NewRxVM creates a new RxVM (RandomX Virtual Machine) using the provided RxDataset and flags.
func NewRxVM(rxDataset *RxDataset, flags ...Flag) (*RxVM, error) {
	if rxDataset.rxCache == nil {
		vm, err := CreateVM(nil, rxDataset.dataset, flags...)
		return &RxVM{
			vm:        vm,
			rxDataset: nil,
		}, err
	}
	vm, err := CreateVM(rxDataset.rxCache.cache, rxDataset.dataset, flags...)
	return &RxVM{
		vm:        vm,
		rxDataset: nil,
	}, err
}

// Close releases the resources associated with the RxVM.
func (vm *RxVM) Close() {
	if vm.vm != nil {
		DestroyVM(vm.vm)
	}
}

// CalcHash calculates and returns the hash of the given input using this RxVM.
func (vm *RxVM) CalcHash(in []byte) []byte {
	return CalculateHash(vm.vm, in)
}

// CalcHashFirst prepares the RxVM to calculate the hash of the given input.
func (vm *RxVM) CalcHashFirst(in []byte) {
	CalculateHashFirst(vm.vm, in)
}

// CalcHashNext calculates and returns the hash of the next input using this RxVM.
func (vm *RxVM) CalcHashNext(in []byte) []byte {
	return CalculateHashNext(vm.vm, in)
}

// UpdateDataset updates the RxVM with a new dataset and cache from the provided RxDataset.
func (vm *RxVM) UpdateDataset(rxDataset *RxDataset) {
	SetVMCache(vm.vm, rxDataset.rxCache.cache)
	SetVMDataset(vm.vm, rxDataset.dataset)
}
