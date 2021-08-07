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
