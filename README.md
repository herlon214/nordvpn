# NordVPN API
This lib uses the nordvpn servers endpoint to fetch data.
```shell script
$ go get github.com/herlon214/nordvpn
```

#### Usage
You can use the server list returned by `FetchData()` but it's also
available many operators to filter, sort the servers. You can even pipe them if needed. 
```go
package main
import (
    "github.com/herlon214/nordvpn"
)

func main() {
    servers, err := nordvpn.FetchData()
    if err != nil {
        // do something
    }

    filteredServers := nordvpn.PipeFilters(
    		nordvpn.FilterOnline(),
    		nordvpn.FilterByCountry("NL"),
    		nordvpn.FilterByTechnology("ikev2"),
    )(servers)
}
```

#### Available operators
|Function|Args example|Description|
|--------|----|-----------|
|FilterOnline(),|-|Filter online servers|
|FilterByTechnology(technologies ...string)|"proxy","ikev2"|Filter by server's technology, you can specify many.|
|FilterByCountry(code string)|"BR"|Filter by country code|
|SortByLoadAsc()|-|Sort servers by load in ascending order|
|SortByLoadDesc|-|Sort servers by load in descending order|

