package main

import (
	"context"
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"

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

	// Initialize dummy records for each asset type
	assetTypes := []externalapi.AssetType{externalapi.KSH, externalapi.KUSD, externalapi.KRV}
	dummyBalances := make(map[string]map[externalapi.AssetType]*pb.AddressBalances)
	for _, assetBalance := range response.AssetBalances {
		for _, addressBalance := range assetBalance.AddressBalances {
			if _, ok := dummyBalances[addressBalance.Address]; !ok {
				dummyBalances[addressBalance.Address] = make(map[externalapi.AssetType]*pb.AddressBalances)
				for _, assetType := range assetTypes {
					dummyBalances[addressBalance.Address][assetType] = &pb.AddressBalances{
						Address:   addressBalance.Address,
						Available: 0,
						Pending:   0,
					}
				}
			}
			dummyBalances[addressBalance.Address][externalapi.AssetTypeFromUint32(assetBalance.AssetType)] = addressBalance
		}
	}

	if conf.Verbose {
		fmt.Printf("%-75s %12s %20s %20s\n", "Address", "Asset Type", "Available", "Pending")
		println("------------------------------------------------------------------------------------------------------------")

		for address, addressBalances := range dummyBalances {
			fmt.Printf("%-75s\n", address)
			for _, assetType := range assetTypes {
				addressBalance := addressBalances[assetType]
				availableStr := utils.FormatKas(addressBalance.Available)
				pendingStr := utils.FormatKas(addressBalance.Pending)
				if addressBalance.Available == 0 {
					availableStr = "--"
				}
				if addressBalance.Pending == 0 {
					pendingStr = "--"
				}
				fmt.Printf("%75s %12s %20s %20s\n", "", assetType.String(), availableStr, pendingStr)
			}
			println("------------------------------------------------------------------------------------------------------------")
		}
	} else {
		totalBalances := make(map[externalapi.AssetType]uint64)
		for _, assetBalance := range response.AssetBalances {
			assetType := externalapi.AssetTypeFromUint32(assetBalance.AssetType)
			for _, addressBalance := range assetBalance.AddressBalances {
				totalBalances[assetType] += addressBalance.Available
			}
		}
		for _, assetType := range assetTypes {
			totalAvailable := totalBalances[assetType]
			availableStr := utils.FormatKas(totalAvailable)
			if totalAvailable == 0 {
				availableStr = "--"
			}
			fmt.Printf("Total balance, %5s %20s\n", assetType.String(), availableStr)
		}
	}

	return nil
}
