package main

import (
	"first_slackbot/slack"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	slack.SendMessageToKoko("hello world")
}
