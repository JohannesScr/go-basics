package my_package

import "testing"

func TestDogYears(t *testing.T) {
	var v int
	v = DogYears(12)
	if v != 84 {
		t.Error("Expected 84 got", v)
	}
}

func TestDogYearsAgain(t *testing.T) {
	var v int
	v = DogYears(12)
	if v != 84 {
		t.Error("Expected 84 got", v)
	}
}

