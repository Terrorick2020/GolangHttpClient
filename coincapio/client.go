package coincapio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero!")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTrip{
				logger: os.Stdout,
				next: http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Assets, nil
}

func (c Client) GetAsset(name string) (Asset, error) {
	if len(name) == 0 {
		return Asset{}, errors.New("name can`t be empty")
	}

	resp, err := c.client.Get(fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name))
	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	var r AssetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return Asset{}, err
	}

	return r.Asset, nil
}