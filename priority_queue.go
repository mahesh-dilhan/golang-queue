package queue

import "container/heap"

type EmergencyCountry struct {
	countryQueue CountryQueue
}

func NewEmergencyCountry() *EmergencyCountry {
	ec := &EmergencyCountry{countryQueue: make(CountryQueue, 0)}
	heap.Init(&ec.countryQueue)
	return ec
}

//helper function
func (ec *EmergencyCountry) registerCountry(country interface{}) {
	potentialCountry := country.(*Country)
	//println(potentialCountry)
	heap.Push(&ec.countryQueue, potentialCountry)
}

func (ec *EmergencyCountry) handleCountry() *Country {
	if ec.countryQueue.Len() == 0 {
		return nil
	}
	//println(ec.countryQueue.Len())
	nextCountry := heap.Pop(&ec.countryQueue)
	return nextCountry.(*Country)
}

func (ec *EmergencyCountry) updateCountryStatus(country Country, severity CountrySeverity) {
	country.severity = severity
	heap.Fix(&ec.countryQueue, country.index)

}
