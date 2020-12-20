package main

import (
	"first_slackbot/slack"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/msg", func(c *gin.Context) {
		text := c.PostForm("msg")
		msg := fmt.Sprintf("<@UTV82PCN7> %s", text )
		slack.SendMessageToKoko(msg)
	})
	r.Run("localhost:8080")
}
