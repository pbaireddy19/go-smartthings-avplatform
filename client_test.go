package smartthings

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Get_Source(t *testing.T) {
	config := Config{&http.Client{}, ""}
	c := Client{&config}

	x, err := c.GetSource("", "")
	fmt.Println(x)
	assert.Equal(t, err, nil)

	/*config := NewConfig("")
	c := Client{config}

	camStreamConfig, err := c.GetSource("", "")

	if err != nil {
		t.Error(err.Error())
	}
	if camStreamConfig != "" {
		var camStmConfig []CameraStreamConfiguration

		err := json.Unmarshal([]byte(camStreamConfig), &camStmConfig)
		if err != nil {
			fmt.Println(fmt.Sprintf("CAM STREAM ERROR \n %v", err.Error()))
		}

		//camStream = append(camStream, camStmConfig[0])
		endpointCapability := EndpointCapabilities{}
		endpointCapabilities := []EndpointCapabilities{}
		endpointCapability.CameraStreamConfigurations = &camStmConfig
		endpointCapabilities = append(endpointCapabilities, endpointCapability)
		x, _ := json.Marshal(endpointCapabilities)
		t.Log("XXX")
		fmt.Println(string(x))
		var xa int
		xa = 5
		fmt.Println(xa)*/
}

/*type EndpointCapabilities struct {
	CameraStreamConfigurations *[]CameraStreamConfiguration `json:"cameraStreamConfigurations,omitempty"`
}

type CameraStreamConfiguration struct {
	Protocols          []string     `json:"protocols"`
	Resolutions        []Resolution `json:"resolutions"`
	AuthorizationTypes []string     `json:"authorizationTypes"`
	VideoCodecs        []string     `json:"videoCodecs"`
	AudioCodecs        []string     `json:"audioCodecs"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}*/
