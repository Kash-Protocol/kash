package main

import (
	"context"
	"fmt"

	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/client"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	if conf.Verbose {
		println("Asset Type    Address                                                                       Available             Pending")
		println("------------------------------------------------------------------------------------------------------------------------")
		for _, assetBalance := range response.AssetBalances {
			assetType := assetBalance.AssetType.String() // Assuming AssetType has a String() method
			for _, addressBalance := range assetBalance.AddressBalances {
				fmt.Printf("%-12s %-75s %-20s %-20s\n", assetType, addressBalance.Address, utils.FormatKas(addressBalance.Available), utils.FormatKas(addressBalance.Pending))
			}
		}
		println("------------------------------------------------------------------------------------------------------------------------")
	} else {
		for _, assetBalance := range response.AssetBalances {
			var totalAvailable, totalPending uint64
			for _, addressBalance := range assetBalance.AddressBalances {
				totalAvailable += addressBalance.Available
				totalPending += addressBalance.Pending
			}
			pendingSuffix := ""
			if totalPending > 0 {
				pendingSuffix = fmt.Sprintf(" (pending %s)", utils.FormatKas(totalPending))
			}
			fmt.Printf("Total balance, %s %s%s\n", assetBalance.AssetType.String(), utils.FormatKas(totalAvailable), pendingSuffix)
		}
	}

	return nil
}
