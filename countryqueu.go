package queue

//priority queue impl
//since we use heap
// sort interface required
type CountryQueue []*Country

func (cq *CountryQueue) Push(patient interface{}) {

}

func (cq *CountrySeverity) Pop() interface{} {
	return nil
}

func (cq *CountryQueue) Len() int {
	return len(*cq)
}

func (cq CountryQueue) swap(a, b int) {
	cq[a], cq[b] = cq[b], cq[a]
	cq[a].index = a
	cq[b].index = b
}

//We need to use map heap based on severity
func (cq CountryQueue) Less(a, b int) bool {
	return cq[a].severity > cq[b].severity
}
