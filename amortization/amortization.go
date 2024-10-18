package amortization

import "math"

type Amortization struct{}

type LoanInformation struct {
	InterestRate float64
	Principle    int32
	Term         int32
}

func (l *LoanInformation) CalculateMonthlyPayment() float64 {
	return calculate(l.InterestRate, l.Term, l.Principle)
}

func (a *Amortization) CalculateMonthlyPayment(interestRate float64, term int32, principle int32) float64 {
	return calculate(interestRate, term, principle)
}

func calculate(interestRate float64, term int32, principle int32) float64 {
	i := interestRate / 12
	t := term * 12
	x := math.Pow(1+i, float64(t))
	numerator := i * x
	denominator := x - 1
	return math.Round(float64(principle)*(numerator/denominator)*100) / 100
}
