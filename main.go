package main

import (
	"amortization-calc-go/pkg/amortization"
	"flag"
	"fmt"
)

func main() {
	p := flag.Int("p", 100000, "principle")
	i := flag.Float64("i", 0.05, "interest rate")
	t := flag.Int("t", 30, "term in years")
	flag.Parse()
	loanInfo := amortization.LoanInformation{Principle: int32(*p), InterestRate: *i, Term: int32(*t)}
	fmt.Println(loanInfo.CalculateMonthlyPayment())
}
