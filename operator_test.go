package nordvpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeFilters(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{
		Name:   "A",
		Status: "online",
		Load:   40,
		Technologies: []Technology{
			{Identifier: "openvpn_udp"},
			{Identifier: "proxy"},
		},
		Locations: []Location{
			{Country: Country{Code: "US"}},
		},
	})
	servers = append(servers, Server{
		Name:   "B",
		Status: "offline",
		Load:   99,
		Technologies: []Technology{
			{Identifier: "proxy_ssl_cybersec"},
		},
		Locations: []Location{
			{Country: Country{Code: "BR"}},
		},
	})
	servers = append(servers, Server{
		Name:   "C",
		Status: "online",
		Load:   70,
		Technologies: []Technology{
			{Identifier: "openvpn_tcp"},
			{Identifier: "proxy_ssl_cybersec"},
		},
		Locations: []Location{
			{Country: Country{Code: "NL"}},
		},
	})
	servers = append(servers, Server{
		Name:   "D",
		Status: "online",
		Load:   20,
		Technologies: []Technology{
			{Identifier: "ikev2"},
			{Identifier: "proxy"},
		},
		Locations: []Location{
			{Country: Country{Code: "NL"}},
		},
	})

	filteredServers := PipeOperators(
		FilterOnline(),
		FilterByCountry("NL"),
		FilterByTechnology("ikev2"),
	)(servers)

	assert.Equal(t, 1, len(filteredServers))
	assert.Equal(t, "D", filteredServers[0].Name)

}

func TestSortByLoadAsc(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{Name: "A", Load: 15})
	servers = append(servers, Server{Name: "B", Load: 5})
	servers = append(servers, Server{Name: "C", Load: 20})
	servers = append(servers, Server{Name: "D", Load: 10})

	sortedAsc := SortByLoadAsc()(servers)
	assert.Equal(t, 4, len(sortedAsc))

	assert.Equal(t, 5, sortedAsc[0].Load)
	assert.Equal(t, 10, sortedAsc[1].Load)
	assert.Equal(t, 15, sortedAsc[2].Load)
	assert.Equal(t, 20, sortedAsc[3].Load)
}

func TestSortByLoadDesc(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{Name: "A", Load: 15})
	servers = append(servers, Server{Name: "B", Load: 5})
	servers = append(servers, Server{Name: "C", Load: 20})
	servers = append(servers, Server{Name: "D", Load: 10})

	sortedDesc := SortByLoadDesc()(servers)
	assert.Equal(t, 4, len(sortedDesc))

	assert.Equal(t, 20, sortedDesc[0].Load)
	assert.Equal(t, 15, sortedDesc[1].Load)
	assert.Equal(t, 10, sortedDesc[2].Load)
	assert.Equal(t, 5, sortedDesc[3].Load)
}

func TestFilterOnline(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{Name: "A", Status: "online"})
	servers = append(servers, Server{Name: "B", Status: "offline"})
	servers = append(servers, Server{Name: "C", Status: "online"})

	onlineServers := FilterOnline()(servers)
	assert.Equal(t, 2, len(onlineServers))

	assert.Equal(t, "A", onlineServers[0].Name)
	assert.Equal(t, "C", onlineServers[1].Name)
}

func TestFilterByCountry(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{Name: "A", Locations: []Location{{Country: Country{Code: "US"}}}})
	servers = append(servers, Server{Name: "B", Locations: []Location{{Country: Country{Code: "BR"}}}})
	servers = append(servers, Server{Name: "C", Locations: []Location{{Country: Country{Code: "NL"}}}})
	servers = append(servers, Server{Name: "D", Locations: []Location{{Country: Country{Code: "BR"}}}})

	filteredServers := FilterByCountry("BR")(servers)
	assert.Equal(t, 2, len(filteredServers))

	assert.Equal(t, "B", filteredServers[0].Name)
	assert.Equal(t, "D", filteredServers[1].Name)
}

func TestFilterByTechnologies(t *testing.T) {
	servers := make([]Server, 0)
	servers = append(servers, Server{Name: "A", Technologies: []Technology{{Identifier: "openvpn_tcp"}, {Identifier: "proxy_ssl"}}})
	servers = append(servers, Server{Name: "B", Technologies: []Technology{{Identifier: "proxy_ssl_cybersec"}, {Identifier: "proxy_ssl"}}})
	servers = append(servers, Server{Name: "C", Technologies: []Technology{{Identifier: "proxy"}, {Identifier: "ikev2"}}})
	servers = append(servers, Server{Name: "D", Technologies: []Technology{{Identifier: "openvpn_udp"}, {Identifier: "proxy_ssl_cybersec"}}})

	filteredServers := FilterByTechnology("proxy_ssl_cybersec", "ikev2")(servers)
	assert.Equal(t, 3, len(filteredServers))

	assert.Equal(t, "B", filteredServers[0].Name)
	assert.Equal(t, "C", filteredServers[1].Name)
	assert.Equal(t, "D", filteredServers[2].Name)
}
