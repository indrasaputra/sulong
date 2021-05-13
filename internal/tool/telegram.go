package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/indrasaputra/sulong/entity"
)

const (
	messageTemplateName = "TaniFund Project"
	messageTemplate     = `
Proyek Baru!

{{.Title}}
Bunga {{.InterestTarget}}%
Dibuka {{.HumanPublishedAt}}
Target {{.TargetFund}}
Tenor {{.Tenor}} bulan
Link {{.ProjectLink}}
`
)

// TelegramMessage ...
type TelegramMessage struct {
	ChatID                int    `json:"chat_id"`
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

// TelegramBot acts as telegram bot.
type TelegramBot struct {
	client          *http.Client
	url             string
	token           string
	messageTemplate *template.Template
}

// NewTelegramBot creates an instance of TelegramBot.
func NewTelegramBot(client *http.Client, url, token string) *TelegramBot {
	return &TelegramBot{
		client:          client,
		url:             url,
		token:           token,
		messageTemplate: template.Must(template.New(messageTemplateName).Parse(messageTemplate)),
	}
}

// Notify notifies the recipient about a project.
func (tb *TelegramBot) Notify(ctx context.Context, recipientID int, project *entity.Project) error {
	var out bytes.Buffer
	if err := tb.messageTemplate.Execute(&out, project); err != nil {
		return err
	}

	message := &TelegramMessage{
		ChatID:                recipientID,
		Text:                  out.String(),
		DisableWebPagePreview: true,
	}

	url := fmt.Sprintf("%s%s/sendMessage", tb.url, tb.token)
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := tb.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}
