## geolocator

[![GoDoc](https://godoc.org/github.com/nathan-osman/geolocator?status.svg)](https://godoc.org/github.com/nathan-osman/geolocator)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an abstract interface for using IP geolocation databases from the following providers:

- [IP2Location™](https://www.ip2location.com/database/ip2location) (including [IP2Location™ Lite](https://lite.ip2location.com/))

### Usage

To use this package in your application, begin by adding the following import:

```go
import "github.com/nathan-osman/geolocator"
```

Providers are included in the subdirectories of this package. For example, to load and query an IP2Location™ database, one would use the following snippet:

```go
import "github.com/nathan-osman/geolocator/ip2location"

provider, err := ip2location.New("path/to/file.bin")
if err != nil {
    // TODO: handle error
}
defer provider.Close()

response, err := provider.Geolocate("123.10.44.12")
if err != nil {
    // TODO: handle error
}

fmt.Printf("Country: %s\n", response.CountryName)
```
