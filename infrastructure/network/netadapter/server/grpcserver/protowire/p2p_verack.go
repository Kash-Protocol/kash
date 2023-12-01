package protowire

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KashdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KashdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *KashdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
