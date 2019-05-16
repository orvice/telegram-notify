package main

import (
	"flag"
	"fmt"
	"github.com/orvice/utils/env"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	telegramToken  string
	telegramChatID int64
	result         string

	bot *tgbotapi.BotAPI
)

func initEnv() {
	telegramToken = env.Get("TELEGRAM_TOKEN")
	telegramChatID = int64(env.GetInt("TELEGRAM_CHATID"))

	result = env.Get("result")
}

func initBot() error {
	var err error
	bot, err = tgbotapi.NewBotAPI(telegramToken)
	return err
}

func main() {
	var err error
	initEnv()

	flag.StringVar(&result,"r","unknown","result")
	flag.Parse()

	err = initBot()
	if err != nil {
		fmt.Println("bot init fail: ", err)
		os.Exit(1)
	}

	format := `
*%s [%s/%s] Build Result*

- 🍊 PipelineUrl: %s
- 🍭 Commit Message: %s

`

	url := fmt.Sprintf("%s/pipelines/%s", Env("CI_PROJECT_URL"), Env("CI_PIPELINE_ID"))
	if len(result) != 0 {

		message := fmt.Sprintf(format, result,
			Env("CI_PROJECT_NAMESPACE"), Env("CI_PROJECT_NAME"), url, Env("CI_COMMIT_TITLE"))

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

func Env(k string) string {
	return os.Getenv(k)
}
