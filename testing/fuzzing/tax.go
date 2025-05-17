package tax

import "time"

func CalculateTax(income float64) float64 {
	if income <= 0 {
		return 0.0
	}
	if income >= 10000 && income < 20000 {
		return 10.0
	}
	if income >= 20000 {
		return 20.0
	}
	return 5.0
}

func CalculateTaxBadly(income float64) float64 {
	time.Sleep(time.Millisecond)
	if income == 0 {
		return 0.0
	}
	if income >= 10000 {
		return 10.0
	}
	return 5.0
}
