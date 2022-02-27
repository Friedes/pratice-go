package cal

import (
	"testing"
)

func TestCase0CalculateTax(t *testing.T) {
	name := "ม่อนด์"

	givenIncome := 150000.0
	givenLtf := 0.0
	givenAssure := 0.0

	wantTax := 4500.0

	getTax := calculateTax(givenIncome, givenLtf, givenAssure)
	if wantTax != getTax {
		t.Errorf("%s\n >>> givenIncome %10.2f, givenLtf %10.2f, givenAssure %10.2f & want %10.2f -----> but %10.2f\n", name, givenIncome, givenLtf, givenAssure, wantTax, getTax)
	}
}

func TestCase1CalculateTax(t *testing.T) {
	name := "ธเนศ"

	givenIncome := 2120000.0
	givenLtf := 300000.0
	givenAssure := 150000.0

	wantTax := 312500.0

	getTax := calculateTax(givenIncome, givenLtf, givenAssure)
	if wantTax != getTax {
		t.Errorf("%s\n >>> givenIncome %10.2f, givenLtf %10.2f, givenAssure %10.2f & want %10.2f -----> but %10.2f\n", name, givenIncome, givenLtf, givenAssure, wantTax, getTax)
	}
}

func TestCase2CalculateTax(t *testing.T) {
	name := "จุฬาลักษณ์"

	givenIncome := 845000.0
	givenLtf := 0.0
	givenAssure := 30000.0

	wantTax := 73500.0

	getTax := calculateTax(givenIncome, givenLtf, givenAssure)
	if wantTax != getTax {
		t.Errorf("%s\n >>> givenIncome %10.2f, givenLtf %10.2f, givenAssure %10.2f & want %10.2f -----> but %10.2f\n", name, givenIncome, givenLtf, givenAssure, wantTax, getTax)
	}
}

func TestCase3CalculateTax(t *testing.T) {
	name := "สุขเกษม"

	givenIncome := 1050000.0
	givenLtf := 25000.0
	givenAssure := 80000.0

	wantTax := 99500.0

	getTax := calculateTax(givenIncome, givenLtf, givenAssure)
	if wantTax != getTax {
		t.Errorf("%s\n >>> givenIncome %10.2f, givenLtf %10.2f, givenAssure %10.2f & want %10.2f -----> but %10.2f\n", name, givenIncome, givenLtf, givenAssure, wantTax, getTax)
	}
}

func TestCase4CalculateTax(t *testing.T) {
	name := "สุพร"

	givenIncome := 5550000.0
	givenLtf := 1200000.0
	givenAssure := 400000.0

	wantTax := 1339000.0

	getTax := calculateTax(givenIncome, givenLtf, givenAssure)
	if wantTax != getTax {
		t.Errorf("%s\n >>> givenIncome %10.2f, givenLtf %10.2f, givenAssure %10.2f & want %10.2f -----> but %10.2f\n", name, givenIncome, givenLtf, givenAssure, wantTax, getTax)
	}
}
