package pow

import (
	"github.com/Kash-Protocol/kashd/util/randomx"
	"github.com/pkg/errors"
	"sync"
)

// RxVMPool represents a pool of RandomX VMs.
type RxVMPool struct {
	vmChan chan *randomx.RxVM
}

var (
	// globalRxVMPool is the global instance of RxVMPool.
	globalRxVMPool *RxVMPool
	once           sync.Once
)

func init() {
	once.Do(func() {
		var err error
		globalRxVMPool, err = NewRxVMPool(2)
		if err != nil {
			panic(errors.Wrap(err, "failed to initialize global RandomX VM pool"))
		}
	})
}

// CalcGlobalVMHash calculates the hash using one of the RandomX VMs from the global pool.
func CalcGlobalVMHash(data []byte) []byte {
	return globalRxVMPool.CalcHash(data)
}

// NewRxVMPool initializes a new pool of RandomX VMs with the given size.
func NewRxVMPool(poolSize int) (*RxVMPool, error) {
	vmChan := make(chan *randomx.RxVM, poolSize)

	for i := 0; i < poolSize; i++ {
		vm, err := createRxVM()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create RandomX VM")
		}
		vmChan <- vm
	}

	return &RxVMPool{vmChan: vmChan}, nil
}

// createRxVM creates a new RandomX VM instance.
func createRxVM() (*randomx.RxVM, error) {
	rxDataset, err := randomx.NewRxDataset(randomx.FlagDefault)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create RandomX dataset")
	}

	vm, err := randomx.NewRxVM(rxDataset, randomx.FlagDefault)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create RandomX VM")
	}

	return vm, nil
}

// CalcHash calculates the hash using one of the available RandomX VMs.
func (p *RxVMPool) CalcHash(data []byte) []byte {
	vm := <-p.vmChan // Get a VM from the pool (non-blocking)
	hash := vm.CalcHash(data)
	p.vmChan <- vm // Return the VM back to the pool
	return hash
}

// Cleanup safely destroys all RandomX VMs in the pool.
func (p *RxVMPool) Cleanup() {
	close(p.vmChan) // Close the channel before cleanup
	for vm := range p.vmChan {
		if vm != nil {
			vm.Close()
		}
	}
}
