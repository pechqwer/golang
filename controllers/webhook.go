package controllers

import (
	"encoding/json"
	"net/http"
	reqApi "servers/restAPI/utils"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
	Postback struct {
		Data string `json:"data"`
	} `json:"postback"`
}

type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func Webhook() gin.HandlerFunc {
	return func(c *gin.Context) {

		channel := c.Request.URL.Query().Get("c")
		if channel == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "channel not found !!",
			})
			return
		}

		// client := &http.Client{}
		// bot, _ := linebot.New("3d9bc22069c513bf432891f005763983", "9TAKxpAvOCF3VuGR5dAcurrZ1mKxg/11ZqenOdhdZ7CvLu3/DSFLbqVP16LtZcbVvi/2xYE5TiUT8cXkIE3LlQ5/uhy3zb5rBhSy1ATch8NLMwl9hYvSLqmzLfg/3xZ1Y+PCvS+5rd3vKY+6kzGRrwdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
		// events, _ := bot.ParseRequest(c.Request)
		// var newTextMessages []linebot.SendingMessage
		// for _, event := range events {
		// 	fmt.Printf("event.Message: %v\n", event.Message)
		// 	if event.Type == linebot.EventTypeMessage {
		// 		userID := event.Source.UserID
		// 		replyToken := event.ReplyToken
		// 		newTextMessages = append(newTextMessages, linebot.NewTextMessage("Display Name: "+userID))
		// 		newTextMessages = append(newTextMessages, linebot.NewTextMessage("Display Name: "+userID))
		// 		_, err := bot.ReplyMessage(replyToken, newTextMessages...).Do()
		// 		if err != nil {
		// 			_, err := bot.PushMessage(userID, newTextMessages...).Do()
		// 			if err != nil {
		// 				println("dsadas")
		// 			}
		// 		}
		// 		println(userID)
		// 	}
		// }

		var body Body
		json.NewDecoder(c.Request.Body).Decode(&body)

		eventsx := body.Events[0]
		typeEvent := eventsx.Type
		message := eventsx.Message

		switch typeEvent {
		case "message":
			switch message.Type {
			case "text":
				result := make(chan *http.Response)
				text := Text{
					Type: "text",
					Text: "ข้อความเข้ามา : " + eventsx.Message.Text + " ยินดีต้อนรับ : ",
				}
				message := ReplyMessage{
					ReplyToken: eventsx.ReplyToken,
					Messages: []Text{
						text,
					},
				}

				value, _ := json.Marshal(message)

				reqApi.SendLineReply("https://api.line.me/v2/bot/message/reply", value, result, "9TAKxpAvOCF3VuGR5dAcurrZ1mKxg/11ZqenOdhdZ7CvLu3/DSFLbqVP16LtZcbVvi/2xYE5TiUT8cXkIE3LlQ5/uhy3zb5rBhSy1ATch8NLMwl9hYvSLqmzLfg/3xZ1Y+PCvS+5rd3vKY+6kzGRrwdB04t89/1O/w1cDnyilFU=")

			}

		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"events": "asdas",
		})
	}
}
