package messages

import (
	"bytes"
	"github.com/canopener/serverlog"
)

// AliasDeniedMessage is the struct which represents an Alias denied message
type AliasDeniedMessage struct {
	MessageType uint8
	Reason      string
}

// NewAliasDeniedMessage returns an instance of AliasDeniedMessage based on params
func NewAliasDeniedMessage(reason string) AliasDeniedMessage {
	return AliasDeniedMessage{
		MessageType: TypeAliasDenied,
		Reason:      reason,
	}
}

// NewAliasDeniedMessageFromBytes returns an instance of AliasDeniedMessage based
// on a slice of bytes
func NewAliasDeniedMessageFromBytes(messageBytes []byte) AliasDeniedMessage {
	message := AliasDeniedMessage{}
	buff := bytes.NewBuffer(messageBytes)
	typeByte, err := buff.ReadByte()
	if err != nil {
		serverlog.Fatal("AliasDenied ", err)
	}

	message.MessageType = uint8(typeByte)
	message.Reason, err = buff.ReadString(NullTerm)
	if err != nil {
		serverlog.Fatal("AliasDenied ", err)
	}

	return message
}

// Bytes returns a slice of bytes representing an AliasDeniedMessage
// which can be sent through a connection
func (ms *AliasDeniedMessage) Bytes() []byte {
	typeBytes := make([]byte, 1)
	typeBytes[0] = byte(ms.MessageType)
	ReasonBytes := append([]byte(ms.Reason), NullTerm)
	return append(typeBytes, ReasonBytes...)
}