package nordvpn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FetchData() ([]Server, error) {
	res, err := http.Get("https://api.nordvpn.com/v1/servers")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var servers []Server
	err = json.Unmarshal(body, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
