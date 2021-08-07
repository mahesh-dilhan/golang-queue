package queue

import "testing"

func testpriorityqueuoperation(t *testing.T) {
	ec := EmergencyCountry{}

	next := ec.handleCountry()
	if next != nil {
		t.Error("expected empty")
	}

	ec.registerCountry(&Country{name: "SG", severity: Mild})
	ec.registerCountry(&Country{name: "USA", severity: Critial})

	next = ec.handleCountry()
	if next.name != "USA" {
		t.Error("expect USA")
	}
}
