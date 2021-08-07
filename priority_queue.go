package queue

import "container/heap"

type EmergencyCountry struct {
	countryQueue CountryQueue
}

func NewEmergencyCountry() *EmergencyCountry {
	er := &EmergencyCountry{countryQueue: make(CountryQueue, 0)}
	heap.Init(&er.countryQueue)
	return er
}

func (er *EmergencyCountry) registerCountry(country interface{}) {
	potentialCountry := country.(*Country)
	heap.Push(&er.countryQueue, potentialCountry)
}
