package geolocator

// Response holds values returned from querying the database.
type Response struct {
	// CountryCode is the 2-character ISO code for the country.
	CountryCode string `json:"country_code"`
	// CountryName is the human-friendly name for the country.
	CountryName string `json:"country_name"`
	// Region is the name of the state / province / area within the country.
	Region string `json:"region"`
	// City is the name of the city within the region or country.
	City string `json:"city"`
	// Latitude is the geographic latitude in degrees.
	Latitude float64 `json:"latitude"`
	// Longitude is the geographic longitude in degrees.
	Longitude float64 `json:"longitude"`
	// Timezone is a timezone name from the IANA timezone database.
	Timezone string `json:"timezone"`
}
