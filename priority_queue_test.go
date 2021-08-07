package queue

import "testing"

func testpriorityqueuoperation(t *testing.T) {
	ec := EmergencyCountry{}

	next := ec.handleCountry()
	if next != nil {
		t.Error("expected empty")
	}
}
