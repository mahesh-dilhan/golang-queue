package queue

//priority queue impl
//since we use heap
// sort interface required
type CountryQueue []*Country

func (cq *CountryQueue) Pop() interface{} {
	return nil
}

func (cq *CountryQueue) Push(countryData interface{}) {
	country := countryData.(*Country)
	country.index = len(*cq) // -> because of heap interface
	*cq = append(*cq, country)
}

func (cq *CountryQueue) Len() int {
	return len(*cq)
}

func (cq CountryQueue) Swap(a, b int) {
	cq[a], cq[b] = cq[b], cq[a]
	cq[a].index = a
	cq[b].index = b
}

//We need to use map heap based on severity
func (cq CountryQueue) Less(a, b int) bool {
	return cq[a].severity > cq[b].severity
}
