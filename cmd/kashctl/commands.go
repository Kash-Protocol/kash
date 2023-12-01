package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Kash-Protocol/kashd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KashdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KashdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KashdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KashdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KashdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KashdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KashdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KashdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KashdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KashdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KashdMessage_BanRequest{}),
	reflect.TypeOf(protowire.KashdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
