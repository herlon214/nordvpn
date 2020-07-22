package nordvpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	_, err := FetchData()
	assert.Nil(t, err)
}
