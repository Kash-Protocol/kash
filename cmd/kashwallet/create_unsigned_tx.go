package main

import (
	"context"
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"os"

	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/client"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/constants"
)

func createUnsignedTransaction(conf *createUnsignedTransactionConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	sendAmountSompi := uint64(conf.SendAmount * constants.SompiPerKash)
	response, err := daemonClient.CreateUnsignedTransactions(ctx, &pb.CreateUnsignedTransactionsRequest{
		From:                     conf.FromAddresses,
		Address:                  conf.ToAddress,
		AssetType:                externalapi.AssetTypeFromString(conf.AssetType).ToUint32(),
		Amount:                   sendAmountSompi,
		IsSendAll:                conf.IsSendAll,
		UseExistingChangeAddress: conf.UseExistingChangeAddress,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Created unsigned transaction")
	fmt.Println(encodeTransactionsToHex(response.UnsignedTransactions))

	return nil
}
