package smartthings

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Get_Source(t *testing.T) {
	config := Config{&http.Client{}, ""}
	c := Client{&config}

	_, err := c.GetSource("")

	assert.Equal(t, err, nil)
}
