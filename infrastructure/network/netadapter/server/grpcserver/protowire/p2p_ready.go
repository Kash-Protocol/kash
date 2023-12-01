package protowire

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KashdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KashdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *KashdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
