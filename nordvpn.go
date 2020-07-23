package nordvpn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type NordVPN struct {
	Servers      []Server
	Operators    []Operator
	MaxCacheTime time.Duration
	UpdatedAt    time.Time

	logger Logger
}

func New(maxCacheTime time.Duration, logger Logger) NordVPN {
	return NordVPN{
		Servers:      make([]Server, 0),
		Operators:    make([]Operator, 0),
		MaxCacheTime: maxCacheTime,
		UpdatedAt:    time.Unix(0, 0),
		logger:       logger,
	}
}

func (n *NordVPN) SetOperators(operators ...Operator) {
	n.Operators = append(n.Operators, operators...)
}

func (n *NordVPN) EnableAutoUpdate() {
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				if n.expired() {
					err := n.UpdateServers()
					if err != nil {
						n.logger.Printf("Failed to auto update NordVPN server list: %s", err.Error())
					}
				}
			}
		}
	}()
}

func (n *NordVPN) UpdateServers() error {
	res, err := http.Get("https://api.nordvpn.com/v1/servers?limit=99999999")
	if err != nil {
		n.logger.Printf("Failed to download NordVPN servers: %s\n", err.Error())
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		n.logger.Printf("Failed to read response body: %s\n", err.Error())
		return err
	}
	defer res.Body.Close()

	var servers []Server
	err = json.Unmarshal(body, &servers)
	if err != nil {
		n.logger.Printf("Failed to unmarshal servers: %s\n", err.Error())
		return err
	}

	// Parse using operators
	if len(n.Operators) > 0 {
		servers = PipeOperators(n.Operators...)(servers)
	}

	n.Servers = servers
	n.UpdatedAt = time.Now()

	return nil
}

func (n *NordVPN) Get() []Server {
	if n.expired() {
		n.UpdateServers()
	}

	return n.Servers
}

// Check if it needs to update the servers
func (n *NordVPN) expired() bool {
	cacheExpired := n.UpdatedAt.Add(n.MaxCacheTime)

	return time.Now().After(cacheExpired)
}
