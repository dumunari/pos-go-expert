package tax

import "errors"

func CalculateTax(income float64) (float64, error) {
	if income <= 0 {
		return 0.0, errors.New("income must be greater than 0")
	}
	if income >= 10000 && income < 20000 {
		return 10.0, nil
	}
	if income >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}
