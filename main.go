package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	// Add the IDs of the channels you want to block
	blockList := []int64{}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Handle incoming messages
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Check if the message is from a blocked channel
		for _, id := range blockList {
			if update.Message.Chat.ID == id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not allowed to access this channel.")
				_, err = bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				break
			}
		}
	}

	// Wait for a signal to terminate the bot
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
}