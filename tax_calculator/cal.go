package cal

const expense float64 = 60000.0

const limitLtf float64 = 200000.0
const limitAssure float64 = 100000.0

func calculateTax(income float64, ltf float64, assure float64) float64 {
	var tax float64

	if limitLtf < ltf {
		ltf = limitLtf
	}

	if limitAssure < assure {
		assure = limitAssure
	}

	net := income - expense - ltf - assure

	// Keep it simple :D
	if net < 150000 {
		tax = net * 0.05
	} else if net <= 300000 {
		tax += (net-150000)*0.05 + 7500
	} else if net <= 500000 {
		tax += (net-300000)*0.1 + 15000
	} else if net <= 750000 {
		tax += (net-500000)*0.15 + 35000
	} else if net <= 1000000 {
		tax += (net-750000)*0.2 + 72500
	} else if net <= 2000000 {
		tax += (net-1000000)*0.25 + 122500
	} else if net <= 5000000 {
		tax += (net-2000000)*0.3 + 372500
	} else {
		tax += (net-5000000)*0.35 + 1272500
	}

	return tax
}
