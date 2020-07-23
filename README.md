# NordVPN API
[![codecov](https://codecov.io/gh/herlon214/nordvpn/branch/master/graph/badge.svg)](https://codecov.io/gh/herlon214/nordvpn)
[![Build Status](https://travis-ci.org/herlon214/nordvpn.svg?branch=master)](https://travis-ci.org/herlon214/nordvpn)

This lib uses the nordvpn servers endpoint to fetch data.
```shell script
$ go get github.com/herlon214/nordvpn
```

#### Usage
```go
package main
import (
    "time"

    "github.com/herlon214/nordvpn"
    "github.com/sirupsen/logrus"
)

func main() {
    maxCacheTime := time.Hour * 1
    logger := logrus.New()
    nvpn := nordvpn.New(maxCacheTime, logger)
    nvpn.SetOperators(
    		nordvpn.FilterOnline(),
    		nordvpn.FilterByCountry("NL"),
    		nordvpn.FilterByTechnology("ikev2"),
    )
    
    nvpn.EnableAutoUpdate() // Optional, will auto update the server list when the cache is expired
    
    // Fetch the servers
    servers := nvpn.Get() // []Servers
    
    // do something with servers
}
```

#### Available operators
|Function|Args example|Description|
|--------|----|-----------|
|FilterOnline(),|-|Filter online servers|
|FilterByTechnology(technologies ...string)|"proxy","ikev2"|Filter by server's technology, you can specify many|
|FilterByCountry(code string)|"BR"|Filter by country code|
|SortByLoadAsc()|-|Sort servers by load in ascending order|
|SortByLoadDesc|-|Sort servers by load in descending order|

