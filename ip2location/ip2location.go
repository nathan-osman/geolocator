package ip2location

import (
	"github.com/ip2location/ip2location-go"
	"github.com/nathan-osman/geolocator"
)

type ip2locationProvider struct {
	db *ip2location.DB
}

// New creates a Provider instance for reading from an ip2location database.
func New(path string) (geolocator.Provider, error) {
	d, err := ip2location.OpenDB(path)
	if err != nil {
		return nil, err
	}
	return &ip2locationProvider{
		db: d,
	}, nil
}

func (i *ip2locationProvider) Geolocate(addr string) (*geolocator.Response, error) {
	r, err := i.db.Get_all(addr)
	if err != nil {
		return nil, err
	}
	return &geolocator.Response{
		CountryCode: r.Country_short,
		CountryName: r.Country_long,
		Region:      r.Region,
		City:        r.City,
		Latitude:    float64(r.Latitude),
		Longitude:   float64(r.Longitude),
		Timezone:    r.Timezone,
	}, nil
}

func (i *ip2locationProvider) Close() {
	i.db.Close()
}
