package tool_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/sulong/entity"
	"github.com/indrasaputra/sulong/internal/tool"
)

const (
	telegramURL         = "http://localhost:8000/telegram"
	telegramToken       = "token"
	telegramRecipientID = 1
)

var (
	invalidTelegramClient = NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusForbidden,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewBufferString(`Forbidden`)),
		}
	})
	validTelegramClient = NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})
)

func TestNewTelegramBot(t *testing.T) {
	t.Run("successfully create an instance of TelegramBot", func(t *testing.T) {
		bot := tool.NewTelegramBot(validTelegramClient, telegramURL, telegramToken)
		assert.NotNil(t, bot)
	})
}

func TestTelegramBot_Notify(t *testing.T) {
	t.Run("fail execute template due to invalid project", func(t *testing.T) {
		bot := tool.NewTelegramBot(validTelegramClient, telegramURL, telegramToken)

		err := bot.Notify(ctx, telegramRecipientID, nil)

		assert.NotNil(t, err)
	})

	t.Run("can't create request due to invalid url", func(t *testing.T) {
		bot := tool.NewTelegramBot(validTelegramClient, "#$%&!*#)ll.com", telegramToken)

		err := bot.Notify(ctx, telegramRecipientID, taniFundProjectEntity())

		assert.NotNil(t, err)
	})

	t.Run("user forbids notification", func(t *testing.T) {
		bot := tool.NewTelegramBot(invalidTelegramClient, telegramURL, telegramToken)

		err := bot.Notify(ctx, telegramRecipientID, taniFundProjectEntity())

		assert.NotNil(t, err)
	})

	t.Run("successfully notify user about a project", func(t *testing.T) {
		bot := tool.NewTelegramBot(validTelegramClient, telegramURL, telegramToken)

		err := bot.Notify(ctx, telegramRecipientID, taniFundProjectEntity())

		assert.Nil(t, err)
	})
}

func taniFundProjectEntity() *entity.Project {
	data := `{"id":"54f0b448-1abf-4720-aba8-d890f13c762d","title":"Budidaya Nanas Malang - 2","pricePerUnit":100000,"maxUnit":2554,"startAt":"2021-05-19T17:00:00Z","endAt":"2021-11-19T16:59:00Z","publishedAt":"2021-05-14T12:00:00Z","cutoffAt":"2021-05-15T12:00:00Z","interestTarget":15,"urlSlug":"malang-s-pineapple-cultivation-2","projectStatus":{"id":5,"description":"Menunggu Fundraising"},"humanPublishedAt":"Saturday, 14-May-21 19:00:00 UTC","humanCutoffAt":"Sunday, 15-May-21 19:00:00 UTC","projectLink":"https://tanifund.com/project/malang-s-pineapple-cultivation-2","targetFund":255400000,"tenor":6}`

	var res entity.Project
	_ = json.Unmarshal([]byte(data), &res)
	return &res
}
