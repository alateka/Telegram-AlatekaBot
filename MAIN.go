package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))

	c := bot.Client()

	bot.HandleMessage("hola", func(m *tbot.Message) {

		fmt.Print("|=> "+m.Chat.FirstName, " (Alias: ", m.Chat.Username, ")", " - Mensaje ==> ", m.Text, "\n")
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)

		time.Sleep(1 * time.Second)

		c.SendMessage(m.Chat.ID, "Hola, como estas "+m.Chat.Username)
	})

	err := bot.Start()

	if err != nil {

		log.Fatal(err)
	}
}
