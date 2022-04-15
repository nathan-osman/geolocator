package geolocator

// Provider is an interface for geolocation providers to implement.
type Provider interface {
	Geolocate(addr string) (*Response, error)
	Close()
}
