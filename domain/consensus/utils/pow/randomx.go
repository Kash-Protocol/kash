package pow

import (
	"github.com/Kash-Protocol/kashd/util/randomx"
	"github.com/pkg/errors"
	"sync"
)

// RxVMPool represents a pool of RandomX VMs.
type RxVMPool struct {
	vmChan chan *randomx.RxVM
	size   int
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

	return &RxVMPool{vmChan: vmChan, size: poolSize}, nil
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

// ResizePool adjusts the size of the RxVMPool to the specified size.
// This function increases the pool size by creating new RandomX VM instances
// and transferring existing ones to a new channel with the desired size.
//
// Note:
//   - This function is not thread-safe and should not be called concurrently.
//   - It does not support shrinking the pool size; attempting to do so will return an error.
//   - The pool can only be resized when it's not in use (i.e., when the current pool size
//     equals the number of VMs in the channel). Attempting to resize while the pool is in use
//     will result in an error.
//
// Parameters:
// - newSize: The desired size of the pool after resizing.
//
// Returns an error if shrinking is attempted, the pool is in use, or if there is an issue
// creating new RandomX VM instances.
func (p *RxVMPool) ResizePool(newSize int) error {
	currentSize := p.size
	if newSize < currentSize {
		return errors.New("shrinking the pool size is not supported")
	}

	if currentSize != len(p.vmChan) {
		return errors.New("cannot resize the pool while it is in use")
	}

	// Create a new channel with the desired size
	newVmChan := make(chan *randomx.RxVM, newSize)

	// Transfer existing VMs to the new channel
	for i := 0; i < currentSize; i++ {
		if vm, ok := <-p.vmChan; ok {
			newVmChan <- vm
		}
	}

	// Calculate the number of VMs to add
	diff := newSize - currentSize

	// Add new VMs to the pool
	for i := 0; i < diff; i++ {
		vm, err := createRxVM()
		if err != nil {
			return errors.Wrap(err, "failed to create RandomX VM during pool resize")
		}
		newVmChan <- vm
	}

	// Replace the old channel with the new one
	p.vmChan = newVmChan
	p.size = newSize

	return nil
}

// ResizeGlobalPool adjusts the size of the global RxVMPool to the specified size.
func ResizeGlobalPool(newSize int) error {
	return globalRxVMPool.ResizePool(newSize)
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
