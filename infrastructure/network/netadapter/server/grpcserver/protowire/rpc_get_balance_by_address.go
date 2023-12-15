package protowire

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KashdMessage_GetBalanceByAddressRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KashdMessage_GetBalanceByAddressRequest is nil")
	}
	return x.GetBalanceByAddressRequest.toAppMessage()
}

func (x *KashdMessage_GetBalanceByAddressRequest) fromAppMessage(message *appmessage.GetBalanceByAddressRequestMessage) error {
	x.GetBalanceByAddressRequest = &GetBalanceByAddressRequestMessage{
		Address: message.Address,
	}
	return nil
}

func (x *GetBalanceByAddressRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBalanceByAddressRequest is nil")
	}
	return &appmessage.GetBalanceByAddressRequestMessage{
		Address: x.Address,
	}, nil
}

func (x *KashdMessage_GetBalanceByAddressResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBalanceByAddressResponse is nil")
	}
	return x.GetBalanceByAddressResponse.toAppMessage()
}

func (x *KashdMessage_GetBalanceByAddressResponse) fromAppMessage(message *appmessage.GetBalanceByAddressResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetBalanceByAddressResponse = &GetBalanceByAddressResponseMessage{
		KshBalance:  message.KSHBalance,
		KusdBalance: message.KUSDBalance,
		KrvBalance:  message.KRVBalance,

		Error: err,
	}
	return nil
}

func (x *GetBalanceByAddressResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBalanceByAddressResponse is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil {
		return &appmessage.GetBalanceByAddressResponseMessage{
			Error: rpcErr,
		}, nil
	}

	return &appmessage.GetBalanceByAddressResponseMessage{
		KSHBalance:  x.KshBalance,
		KUSDBalance: x.KusdBalance,
		KRVBalance:  x.KrvBalance,

		Error: rpcErr,
	}, nil
}
