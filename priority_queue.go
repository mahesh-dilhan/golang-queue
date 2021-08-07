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
	heap.Push(&ec.countryQueue, potentialCountry)
}

func (ec *EmergencyCountry) handleCountry() *Country {
	if ec.countryQueue.Len() == 0 {
		return nil
	}
	return heap.Pop(&ec.countryQueue).(*Country)
}
