package main

import (
	"errors"
	"log"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getVideoId(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}
	if u.Host != "www.youtube.com" {
		return "", errors.New("invalid url! Bot supports only www.youtube.com")
	}
	videoId := u.Query().Get("v")
	if videoId == "" {
		return "", errors.New("invalid url! url must look like this: https://www.youtube.com/watch?v=<video-id>")
	}

	return videoId, nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		_, err := getVideoId(update.Message.Text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		} else {

		}
	}
}
