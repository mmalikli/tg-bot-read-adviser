package main

import (
	"flag"
	"log"
	"tg-bot-read-adviser/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

		// if t, err := token(); err != nil {
	// 	panic
	// }

	tgClient = telegram.New(tgBotHost, MustToken())

	// fetcher = fethcer.New()

	// processort = processor.New()

	// consumer.Start(fetcher, processor)
}

func MustToken() (string) {
	token := flag.String(
		"token-bot-token", 
		"", 
		"token for accessing the tg bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}