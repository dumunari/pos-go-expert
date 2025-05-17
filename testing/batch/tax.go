package tax

func CalculateTax(income float64) float64 {
	if income >= 10000 {
		return 10.0
	}
	return 5.0
}
