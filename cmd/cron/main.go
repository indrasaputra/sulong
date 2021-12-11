package main

import (
	"fmt"
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

	go checkAndNotify(cfg)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
}

func checkAndNotify(cfg *config.Config) {
	client := &http.Client{
		Timeout: timeout,
	}

	getter := tool.NewTaniFundCrawler(client, cfg.TaniFund.URL)
	notifier := tool.NewTelegramBot(client, cfg.Telegram.URL, cfg.Telegram.Token)

	checker := usecase.NewTaniFundProjectChecker(getter, notifier, cfg.Telegram.RecipientID)

	if err := checker.CheckAndNotify(); err != nil {
		log.Printf("check and notify: %v", err)
	} else {
		log.Printf("success check and notify")
	}
}
