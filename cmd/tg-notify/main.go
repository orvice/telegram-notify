package main

import (
	"fmt"
	"github.com/orvice/utils/env"
	"gopkg.in/telegram-bot-api.v4"
	"os"
)

var (
	telegramToken  string
	telegramChatID int64

	bot *tgbotapi.BotAPI
)

func initEnv() {
	telegramToken = env.Get("TELEGRAM_TOKEN")
	telegramChatID = int64(env.GetInt("TELEGRAM_CHATID"))
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
		fmt.Println(err)
		os.Exit(1)
	}

}
