package nordvpn

type Timestamp struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Server struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Station      string       `json:"station"`
	Hostname     string       `json:"hostname"`
	Load         int          `json:"load"`
	Status       string       `json:"status"`
	Locations    []Location   `json:"locations"`
	Services     []Service    `json:"serivces"`
	Technologies []Technology `json:"technologies"`
	Timestamp
}

type Technology struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Pivot      struct {
		Status string `json:"status"`
	} `json:"pivot"`
	Timestamp
}

type Service struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Timestamp
}

type Location struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   Country `json:"country"`
	Timestamp
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"` // 2 Letters code
}

type City struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DnsName   string  `json:"dns_name"`
}
