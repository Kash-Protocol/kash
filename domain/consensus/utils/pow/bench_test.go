package pow

import (
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"sync"
	"testing"
	"time"
)

// TestCalcRandomxHashRate tests the efficiency of CalcHash function using multiple goroutines.
// It calculates the total number of hashes computed per second by all goroutines.
func TestCalcRandomxHashRate(t *testing.T) {
	const (
		testDuration  = 5 * time.Second
		numGoroutines = 2 // Number of goroutines to use
	)
	data := []byte("test data for hashing")

	var (
		hashCount int
		wg        sync.WaitGroup
	)

	wg.Add(numGoroutines)
	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				if time.Since(startTime) >= testDuration {
					return
				}
				_ = CalcGlobalVMHash(data)
				hashCount++
			}
		}()
	}

	wg.Wait()
	hashRate := float64(hashCount) / testDuration.Seconds()
	fmt.Printf("Randomx - Total Hashes per second: %f\n", hashRate)
}

// TestCalcKheavyHashRate tests the efficiency of HeavyHash function using multiple goroutines.
// It calculates the total number of hashes computed per second by all goroutines.
func TestCalcKheavyHashRate(t *testing.T) {
	const (
		testDuration  = 5 * time.Second
		numGoroutines = 5 // Number of goroutines to use
	)
	mockHash := externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
		0x41, 0x1f, 0x8c, 0xd2, 0x6f, 0x3d, 0x41, 0xae,
		0xa3, 0x9e, 0x78, 0x57, 0x39, 0x27, 0xda, 0x24,
		0xd2, 0x39, 0x95, 0x70, 0x5b, 0x57, 0x9f, 0x30,
		0x95, 0x9b, 0x91, 0x27, 0xe9, 0x6b, 0x79, 0xe3,
	})
	mat := generateMatrix(mockHash)

	var (
		heavyHashCount int
		wg             sync.WaitGroup
	)

	wg.Add(numGoroutines)
	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				if time.Since(startTime) >= testDuration {
					return
				}
				_ = mat.HeavyHash(mockHash)
				heavyHashCount++
			}
		}()
	}

	wg.Wait()
	heavyHashRate := float64(heavyHashCount) / testDuration.Seconds()
	fmt.Printf("HeavyHash - Total Hashes per second: %f\n", heavyHashRate)
}
