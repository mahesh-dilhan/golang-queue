package queue

import "testing"

func testenqueandDequeue(t *testing.T) {
	q := &Queue{}

	q.enqueue("SG")
	q.enqueue("USA")

	item1 := q.dequeue()

	if item1 != "SG" {
		t.Error("First item should be SG")
	}
	item2 := q.dequeue()
	if item2 != "USA" {
		t.Error("First item should be USA")
	}
}
