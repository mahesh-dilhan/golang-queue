package queue

type EmergencyCountry struct {
	countryQueue CountryQueue
}

func NewEmergencyCountry() *EmergencyCountry {
	err := &EmergencyCountry{countryQueue: make(CountryQueue, 0)}
}
