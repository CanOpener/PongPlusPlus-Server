package messages

import (
	"bytes"
	"github.com/canopener/serverlog"
)

// CreateGameDeniedMessage is a structure representing a Create game denied message
type CreateGameDeniedMessage struct {
	MessageType uint8
	GameName    string
	Reason      string
}

// NewCreateGameDeniedMessage returns an instance of CreateGameDeniedMessage from params
func NewCreateGameDeniedMessage(gameName, reason string) CreateGameDeniedMessage {
	return CreateGameDeniedMessage{
		MessageType: TypeCreateGameDenied,
		GameName:    gameName,
		Reason:      reason,
	}
}

// NewCreateGameDeniedMessageFromBytes returns an instance of CreateGameDeniedMessage
// from a slice of bytes
func NewCreateGameDeniedMessageFromBytes(messageBytes []byte) CreateGameDeniedMessage {
	message := CreateGameDeniedMessage{}
	buff := bytes.NewBuffer(messageBytes)
	typeByte, err := buff.ReadByte()
	if err != nil {
		serverlog.Fatal("CreateGameDenied ", err)
	}

	message.MessageType = uint8(typeByte)
	message.GameName, err = buff.ReadString(NullTerm)
	if err != nil {
		serverlog.Fatal("CreateGameDenied ", err)
	}
	message.Reason, err = buff.ReadString(NullTerm)
	if err != nil {
		serverlog.Fatal("CreateGameDenied ", err)
	}

	return message
}

// Bytes returns a slice of bytes representing an CreateGameDeniedMessage
// which can be sent through a connection
func (ms *CreateGameDeniedMessage) Bytes() []byte {
	typeBytes := make([]byte, 1)
	typeBytes[0] = byte(ms.MessageType)
	gameNameBytes := append([]byte(ms.GameName), NullTerm)
	reasonBytes := append([]byte(ms.Reason), NullTerm)

	message := append(typeBytes, gameNameBytes...)
	message = append(message, reasonBytes...)
	return message
}
