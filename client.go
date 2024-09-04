package mapon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL:    "https://mapon.com/api/v1",
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) GetUnitDataHistoryPoint(unitID int, datetime string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/unit_data/history_point.json?key=%s&unit_id=%d&datetime=%s", c.BaseURL, c.APIKey, unitID, datetime)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
