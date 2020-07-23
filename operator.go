package nordvpn

import (
	"fmt"
	"sort"
	"strings"
)

type Operator = func(servers []Server) []Server

func PipeOperators(operators ...Operator) Operator {
	return func(servers []Server) []Server {
		for _, filter := range operators {
			servers = filter(servers)
		}

		return servers
	}
}

func FilterOnline() Operator {
	return func(servers []Server) []Server {
		online := make([]Server, 0)

		for _, server := range servers {
			if server.Status == "online" {
				online = append(online, server)
			}
		}

		return online
	}

}

func FilterByTechnology(technologies ...string) Operator {
	// Avoid looping
	allTech := strings.Join(technologies, "|")
	allTech = fmt.Sprintf("|%s|", allTech)

	return func(servers []Server) []Server {
		filtered := make([]Server, 0)

		for _, server := range servers {
			for _, tech := range server.Technologies {
				if strings.Contains(allTech, fmt.Sprintf("|%s|", tech.Identifier)) {
					filtered = append(filtered, server)
				}

			}
		}

		return filtered
	}
}

func FilterByCountry(code string) Operator {
	return func(servers []Server) []Server {
		filtered := make([]Server, 0)

		for _, server := range servers {
			for _, location := range server.Locations {
				if location.Country.Code == code {
					filtered = append(filtered, server)
				}
			}
		}

		return filtered
	}
}

func SortByLoadAsc() Operator {
	return func(servers []Server) []Server {
		sort.SliceStable(servers, func(i int, j int) bool {

			return servers[i].Load < servers[j].Load
		})

		return servers
	}

}

func SortByLoadDesc() Operator {
	return func(servers []Server) []Server {
		sort.SliceStable(servers, func(i int, j int) bool {
			return servers[i].Load > servers[j].Load

		})

		return servers
	}

}
