package main

import (
	"log"
	"net/http"
	"time"

	"github.com/indrasaputra/sulong/internal/config"
	"github.com/indrasaputra/sulong/internal/tool"
	"github.com/indrasaputra/sulong/usecase"
)

const (
	timeout = 5 * time.Second
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Printf("fail initiate config: %v", err)
		return
	}

	client := &http.Client{
		Timeout: timeout,
	}

	getter := tool.NewTaniFundCrawler(client, cfg.TaniFund.URL)
	notifier := tool.NewTelegramBot(client, cfg.Telegram.URL, cfg.Telegram.Token)

	checker := usecase.NewTaniFundProjectChecker(getter, notifier, cfg.Telegram.RecipientID)

	for {
		if err := checker.CheckAndNotify(); err != nil {
			log.Printf("check and notify: %v", err)
		} else {
			log.Printf("success check and notify")
		}
		time.Sleep(1 * time.Hour)
	}
}
