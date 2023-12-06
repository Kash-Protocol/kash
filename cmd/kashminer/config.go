package main

import (
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/pow"
	"github.com/Kash-Protocol/kashd/infrastructure/config"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Kash-Protocol/kashd/util"
	"github.com/pkg/errors"

	"github.com/Kash-Protocol/kashd/version"
	"github.com/jessevdk/go-flags"
)

const (
	defaultLogFilename          = "kashminer.log"
	defaultErrLogFilename       = "kashminer_err.log"
	defaultTargetBlockRateRatio = 2.0
)

var (
	// Default configuration options
	defaultAppDir     = util.AppDir("kashminer", false)
	defaultLogFile    = filepath.Join(defaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(defaultAppDir, defaultErrLogFilename)
	defaultRPCServer  = "localhost"
	defaultWorkers    = 2
)

type configFlags struct {
	ShowVersion           bool     `short:"V" long:"version" description:"Display version information and exit"`
	RPCServer             string   `short:"s" long:"rpcserver" description:"RPC server to connect to"`
	MiningAddr            string   `long:"miningaddr" description:"Address to mine to"`
	NumberOfBlocks        uint64   `short:"n" long:"numblocks" description:"Number of blocks to mine. If omitted, will mine until the process is interrupted."`
	MineWhenNotSynced     bool     `long:"mine-when-not-synced" description:"Mine even if the node is not synced with the rest of the network."`
	Profile               string   `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	TargetBlocksPerSecond *float64 `long:"target-blocks-per-second" description:"Sets a maximum block rate. 0 means no limit (The default one is 2 * target network block rate)"`
	Workers               int      `long:"workers" description:"Number of concurrent mining workers"`
	config.NetworkFlags
}

func parseConfig() (*configFlags, error) {
	cfg := &configFlags{
		RPCServer: defaultRPCServer,
		Workers:   defaultWorkers,
	}
	parser := flags.NewParser(cfg, flags.PrintErrors|flags.HelpFlag)
	_, err := parser.Parse()

	// Show the version and exit if the version flag was specified.
	if cfg.ShowVersion {
		appName := filepath.Base(os.Args[0])
		appName = strings.TrimSuffix(appName, filepath.Ext(appName))
		fmt.Println(appName, "version", version.Version())
		os.Exit(0)
	}

	if err != nil {
		return nil, err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return nil, err
	}

	if cfg.TargetBlocksPerSecond == nil {
		targetBlocksPerSecond := defaultTargetBlockRateRatio / cfg.NetParams().TargetTimePerBlock.Seconds()
		cfg.TargetBlocksPerSecond = &targetBlocksPerSecond
	}

	if cfg.Profile != "" {
		profilePort, err := strconv.Atoi(cfg.Profile)
		if err != nil || profilePort < 1024 || profilePort > 65535 {
			return nil, errors.New("The profile port must be between 1024 and 65535")
		}
	}

	// Currently mainnet mining is not supported
	if cfg.ActiveNetParams.Name == "kash-mainnet" {
		return nil, errors.New("Currently mining is not supported on mainnet")
	}

	err = pow.ResizeGlobalPool(cfg.Workers)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to resize global pool")
	}

	if cfg.MiningAddr == "" {
		return nil, errors.New("--miningaddr is required")
	}

	initLog(defaultLogFile, defaultErrLogFile)

	return cfg, nil
}
