package main

import (
	"fmt"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/orvice/utils/env"
)

var (
	telegramToken  string
	telegramChatID int64
	message        string

	bot *tgbotapi.BotAPI
)

func initEnv() {
	telegramToken = env.Get("TELEGRAM_TOKEN")
	telegramChatID = int64(env.GetInt("TELEGRAM_CHATID"))

	message = env.Get("MESSAGE")
}

func initBot() error {
	var err error
	bot, err = tgbotapi.NewBotAPI(telegramToken)
	return err
}

func main() {
	var err error
	initEnv()
	err = initBot()
	if err != nil {
		fmt.Println("bot init fail: ", err)
		os.Exit(1)
	}

	if len(message) != 0 {
		msg := tgbotapi.NewMessage(telegramChatID, message)
		msg.ParseMode = tgbotapi.ModeMarkdown
		reply, err := bot.Send(msg)
		if err != nil {
			fmt.Println("send error :", err)
			return
		}
		fmt.Println("reply id:  ", reply.MessageID)
	}
}
