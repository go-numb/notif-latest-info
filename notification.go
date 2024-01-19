package notif_latest_info

import (
	"errors"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// ToLine 通知関数本体
func ToLine(to, msg string) error {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		return errors.Join(err, errors.New("make client for line"))
	}

	bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: to,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: msg,
				},
			},
		},
		"", // x-line-retry-key
	)

	return nil
}
