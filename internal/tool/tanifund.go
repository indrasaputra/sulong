package tool

import (
	"context"
	"encoding/json"
	"fmt"
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
	return result.Data.Items, nil
}
