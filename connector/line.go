package connector

import (
	"fmt"
	"gomoka-cryptobot/config"
	"gomoka-cryptobot/core"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// LineConnector Line resolve connector
type LineConnector struct {
	*core.MessengerConnector
}

type textMessage struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}
type eventSource struct {
	UserID  string `json:"userId"`
	Type    string `json:"type"`
	GroupID string `json:"groupId"`
}
type event struct {
	Source     eventSource       `json:"source"`
	ReplyToken string            `json:"replyToken"`
	Message    map[string]string `json:"message"`
	Timestamp  int64             `json:"timestamp"`
}
type lineRequestBody struct {
	Events []event `json:"events"`
}

type ResponsePayload struct {
}

func (line *LineConnector) convertRequestToIncommingMessage(c *gin.Context) core.IncommingMessage {
	return core.IncommingMessage{
		Type: core.TEXT,
	}
}

// RequestHandler for Line provider
func (line *LineConnector) RequestHandler(c *gin.Context) {
	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	// convert payload
	var body lineRequestBody
	err = c.BindJSON(&body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// iterate to resolve message
	for _, event := range body.Events {
		// filter only text message
		if event.Message["type"] == "text" {
			var (
				message  textMessage
				senderID string
			)
			err = mapstructure.Decode(event.Message, &message)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}

			switch event.Source.Type {
			case "group":
				{
					senderID = event.Source.GroupID
				}
			case "user":
				{
					senderID = event.Source.UserID
				}
			}

			fmt.Println(fmt.Sprintf("[Line] %s say: %s", senderID, message.Text))
		}
	}
	c.JSON(http.StatusOK, gin.H{})

	// message := line.convertRequestToIncommingMessage(c)

	// responseMessages := line.GetResponseMessages(message)

	// return result to provider

}

func CreateLineConnector(config config.Config) *LineConnector {

	connector := core.MessengerConnector{
		Config:       config,
		ProviderName: "Line",
	}
	line := &LineConnector{&connector}
	return line
}
