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
	ec.registerCountry(&Country{name: "IND", severity: Critial})

	next := ec.handleCountry()
	if next.name != "USA" {
		t.Error("expect USA", next.name)
	}
}

func TestPriorityqueuoperation3(t *testing.T) {
	ec := NewEmergencyCountry()
	sg := &Country{name: "SG", severity: Mild}
	ec.registerCountry(sg)

	usa := &Country{name: "USA", severity: Mild}
	ec.registerCountry(usa)

	ind := &Country{name: "IND", severity: Mild}
	ec.registerCountry(ind)

	ec.updateCountryStatus(sg, Critial)

	next := ec.handleCountry()
	if next.name != "SG" {
		t.Error("expect SG", next.name)
	}
}
