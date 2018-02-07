package core

import (
	"gomoka-cryptobot/config"

	"github.com/gin-gonic/gin"
)

// MessageType Type of response message, use internally
type MessageType int

const (
	// TEXT message is plain text
	TEXT MessageType = iota
)

// ResponseMessage response messsage after
// resolve command
type ResponseMessage struct {
	Type    MessageType
	Content string
	Target  string
}

// IncommingMessage every connector should parse
// request message to Incomming message
type IncommingMessage struct {
	RawText    string
	FromUserID string
	Type       MessageType
}

// MessengerConnector root abstruct of messenger provider
// include how to resolve message
type MessengerConnector struct {
	Config         config.Config
	ProviderName   string
	RequestHandler gin.HandlerFunc
}

func (m MessengerConnector) GetResponseMessages(message IncommingMessage) []ResponseMessage {

	// resolve data from incomming message

	return []ResponseMessage{
		ResponseMessage{
			Type:    TEXT,
			Content: "Hello World",
			Target:  "Someone",
		},
	}
}

// RequestHandler resolve incommingMessage to ResponseMessage
