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

	x, err := c.GetSource("", "sdsd")
	fmt.Println(x)
	assert.Equal(t, err, nil)
}
