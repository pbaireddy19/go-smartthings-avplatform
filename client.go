package smartthings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pbaireddy19/go-smartthings-avplatform/models"
)

// Client implements a Smartthings client
type Client struct {
	*Config
}

//New client
func New(config *Config) *Client {
	c := &Client{Config: config}

	return c
}

//GetSource - Fetches the source details
func (c *Client) GetSource(sourceID string) (models.GetSourceResponse, error) {
	sourceResp := models.GetSourceResponse{}
	req, err := http.NewRequest("GET", ST_AV_URL+"/source", nil)

	if err != nil {
		return sourceResp, err
	}

	params := url.Values{}
	params.Add("source_id", sourceID)
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Authorization", "Bearer "+c.Config.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Config.HTTPClient.Do(req)

	if err != nil {
		return sourceResp, err
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return sourceResp, err
		}
		fmt.Println(string(body))
		err = json.Unmarshal(body, &sourceResp)

		if err != nil {
			return sourceResp, err
		}
	} else {
		return sourceResp, errors.New("failed to retrieve source info " + resp.Status)
	}

	return sourceResp, err
}
