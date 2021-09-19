package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

func init() { // Load .env

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	alatekaBot := tbot.New(os.Getenv("TELEGRAM_TOKEN")) // Get token telegram

	alatekaClient := alatekaBot.Client()

	alatekaBot.HandleMessage("hola", func(m *tbot.Message) { // Getting message & progress it

		fmt.Print("|=> "+m.Chat.FirstName, " (Alias: ", m.Chat.Username, ")", " - Mensaje ==> ", m.Text, "\n")
		alatekaClient.SendChatAction(m.Chat.ID, tbot.ActionTyping)

		time.Sleep(1 * time.Second)

		alatekaClient.SendMessage(m.Chat.ID, "Hola, como estas "+m.Chat.Username)
	})

	err := alatekaBot.Start()

	if err != nil { // If return a error, show it

		log.Fatal(err)
	}
}
