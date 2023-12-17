package protowire

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KashdMessage_GetCoinSupplyRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.GetCoinSupplyRequestMessage{}, nil
}

func (x *KashdMessage_GetCoinSupplyRequest) fromAppMessage(_ *appmessage.GetCoinSupplyRequestMessage) error {
	x.GetCoinSupplyRequest = &GetCoinSupplyRequestMessage{}
	return nil
}

func (x *KashdMessage_GetCoinSupplyResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KashdMessage_GetCoinSupplyResponse is nil")
	}
	return x.GetCoinSupplyResponse.toAppMessage()
}

func (x *KashdMessage_GetCoinSupplyResponse) fromAppMessage(message *appmessage.GetCoinSupplyResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetCoinSupplyResponse = &GetCoinSupplyResponseMessage{
		MaxSompi:             message.MaxSompi,
		KshCirculatingSompi:  message.KSHCirculatingSompi,
		KusdCirculatingSompi: message.KUSDCirculatingSompi,
		KrvCirculatingSompi:  message.KRVCirculatingSompi,

		Error: err,
	}
	return nil
}

func (x *GetCoinSupplyResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetCoinSupplyResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	return &appmessage.GetCoinSupplyResponseMessage{
		MaxSompi:             x.MaxSompi,
		KSHCirculatingSompi:  x.KshCirculatingSompi,
		KUSDCirculatingSompi: x.KusdCirculatingSompi,
		KRVCirculatingSompi:  x.KrvCirculatingSompi,

		Error: rpcErr,
	}, nil
}
