package tax

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTaxAndSave(income float64, repository Repository) error {
	tax := CalculateTax(income)
	return repository.SaveTax(tax)
}

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
