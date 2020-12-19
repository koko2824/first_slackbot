package slack

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func sendMessage(text string, url string, method string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	message := map[string]string{
		"text": text,
	}

	by, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	js := bytes.NewBuffer(by)
	req, err := http.NewRequest(method, url, js)
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

func SendMessageToKoko(text string) {
	method := "POST"
	url := os.Getenv("SLACK_CH_DEF")
	sendMessage(text, url, method)
}
