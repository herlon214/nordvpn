package nordvpn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func TestAPI(t *testing.T) {
	nvpn := New(time.Second*10, logrus.New())
	nvpn.SetOperators(FilterOnline(), FilterByCountry("BR"), SortByLoadAsc())

	err := nvpn.UpdateServers()
	assert.Nil(t, err)

	servers := nvpn.Get()
	assert.Greater(t, len(servers), 0)
}
