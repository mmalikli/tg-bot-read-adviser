package telegram

import "tg-bot-read-adviser/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client)
