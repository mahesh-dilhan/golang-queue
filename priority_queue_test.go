package queue

import "testing"

func TestPriorityqueuoperation(t *testing.T) {
	ec := NewEmergencyCountry()

	next := ec.handleCountry()
	if next != nil {
		t.Error("expected empty")
	}

}

func TestPriorityqueuoperation2(t *testing.T) {
	ec := NewEmergencyCountry()
	ec.registerCountry(&Country{name: "SG", severity: Mild})
	ec.registerCountry(&Country{name: "USA", severity: Critial})

	next := ec.handleCountry()
	if next.name != "USA" {
		t.Error("expect USA")
	}
}
