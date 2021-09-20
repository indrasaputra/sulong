package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/indrasaputra/sulong/entity"
)

// TaniFundCrawler acts as crawler for getting data from TaniFund.
type TaniFundCrawler struct {
	client *http.Client
	url    string
}

// NewTaniFundCrawler creates an instance of TaniFundCrawler.
func NewTaniFundCrawler(client *http.Client, url string) *TaniFundCrawler {
	return &TaniFundCrawler{
		client: client,
		url:    url,
	}
}

// GetNewestProjects gets newest projects in TaniFund.
func (tfc *TaniFundCrawler) GetNewestProjects(ctx context.Context, numberOfProject int) ([]*entity.Project, error) {
	url := fmt.Sprintf("%s?page=1&itemsPerPage=%d&sort=-cutoffAt", tfc.url, numberOfProject)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []*entity.Project{}, err
	}
	req.Header.Set("Authority", "tanifund.com")
	req.Header.Set("Referer", "https://tanifund.com/projects")
	req.Header.Set("Accept-Language", "id")
	req.Header.Set("Accept", "application/json, text/plain, */*'")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36'")

	resp, err := tfc.client.Do(req)
	if err != nil {
		return []*entity.Project{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var result entity.TaniFund
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []*entity.Project{}, err
	}

	if result.Data == nil {
		log.Printf("got nil data for request: %s\n", url)
		return []*entity.Project{}, nil
	}
	return result.Data.Items, nil
}
