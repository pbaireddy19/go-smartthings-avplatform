package smartthings

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Get_Source(t *testing.T) {
	config := Config{&http.Client{}, "0ba1dc03-a683-4f25-8b85-a275506edb61"}
	c := Client{&config}

	x, err := c.GetSource("c0b92631-e8ab-43c7-a699-2913552c9a6f")
	fmt.Println(x)
	assert.Equal(t, err, nil)
}
