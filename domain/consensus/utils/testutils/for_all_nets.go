package testutils

import (
	"github.com/Kash-Protocol/kashd/domain/consensus"
	// Importing for the side effect of initializing the RandomX VM in the pow package
	_ "github.com/Kash-Protocol/kashd/domain/consensus/utils/pow"
	"github.com/Kash-Protocol/kashd/domain/dagconfig"
	"testing"
)

// ForAllNets runs the passed testFunc with all available networks
// if setDifficultyToMinumum = true - will modify the net params to have minimal difficulty, like in SimNet
func ForAllNets(t *testing.T, skipPow bool, testFunc func(*testing.T, *consensus.Config)) {
	allParams := []dagconfig.Params{
		dagconfig.MainnetParams,
		dagconfig.TestnetParams,
		dagconfig.SimnetParams,
		dagconfig.DevnetParams,
	}

	for _, params := range allParams {
		consensusConfig := consensus.Config{Params: params}
		t.Run(consensusConfig.Name, func(t *testing.T) {
			t.Parallel()
			consensusConfig.SkipProofOfWork = skipPow
			t.Logf("Running test for %s", consensusConfig.Name)
			testFunc(t, &consensusConfig)
		})
	}
}
