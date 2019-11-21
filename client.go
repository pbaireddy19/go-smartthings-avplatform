package smartthings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pbaireddy19/go-smartthings-avplatform/constants"
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
func (c *Client) GetSource(sourceID string, manufacturer string) (string, error) {

	switch manufacturer {
	case constants.CameraSmartthings, constants.CameraUnknown:
		return c.getSmartthingsCamInfo(sourceID)
	default:
		return "", errors.New("Not Found")
	}
	return "", errors.New("Error retrieving the source")
}

func (c *Client) getSmartthingsCamInfo(sourceID string) (string, error) {
	sourceResp := models.GetSourceResponse{}

	req, err := http.NewRequest("GET", ST_AV_URL+"/source", nil)

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("source_id", sourceID)
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Authorization", "Bearer "+c.Config.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Config.HTTPClient.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		//fmt.Println(string(body))
		err = json.Unmarshal(body, &sourceResp)

		if err == nil {
			highProfile := sourceResp.Source.Profiles.High
			mediumProfile := sourceResp.Source.Profiles.Medium
			lowProfile := sourceResp.Source.Profiles.Low

			camConfig := models.CameraConfig{}
			//TODO: Figure out to fetch it dynamically
			camConfig.Protocols = []string{"RTSP"}
			if len(highProfile.Media.Tracks) == 2 {
				if highProfile.Media.Tracks[0].Type == "video" {
					camConfig.Resolutions = []models.Resolution{models.Resolution{Width: highProfile.Media.Tracks[0].Width, Height: highProfile.Media.Tracks[0].Height},
						models.Resolution{Width: mediumProfile.Media.Tracks[0].Width, Height: mediumProfile.Media.Tracks[0].Height},
						models.Resolution{Width: lowProfile.Media.Tracks[0].Width, Height: lowProfile.Media.Tracks[0].Height}}
					camConfig.VideoCodecs = []string{highProfile.Media.Tracks[0].Codec}
					camConfig.AudioCodecs = []string{highProfile.Media.Tracks[1].Codec}
				} else {
					camConfig.Resolutions = []models.Resolution{models.Resolution{Width: highProfile.Media.Tracks[1].Width, Height: highProfile.Media.Tracks[1].Height},
						models.Resolution{Width: mediumProfile.Media.Tracks[1].Width, Height: mediumProfile.Media.Tracks[1].Height},
						models.Resolution{Width: lowProfile.Media.Tracks[1].Width, Height: lowProfile.Media.Tracks[1].Height}}
					camConfig.VideoCodecs = []string{highProfile.Media.Tracks[1].Codec}
					camConfig.AudioCodecs = []string{highProfile.Media.Tracks[0].Codec}
				}

				camConfig.AuthorizationTypes = []string{"BASIC"}

				responseMarshalled, err := json.Marshal([]models.CameraConfig{camConfig})

				if err != nil {
					return "", err
				}

				return string(responseMarshalled), err
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		return "", errors.New("failed to retrieve source info " + resp.Status)
	}

	return "", err
}
