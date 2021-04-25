package cash_flows

import "fmt"

type BondCalculator struct {
	Institution string
	Numperiods  int64
	Principle   float64
	CouponValue float64
}

// calculate the interest rate
func (bc *BondCalculator) InterestRate() float64 {
	return bc.CouponValue / bc.Principle
}

// function to format print statements
func (bc *BondCalculator) String() string {
	interestRate := bc.InterestRate()
	return fmt.Sprintf("%s: at %.2f for %2d", bc.Institution, interestRate, bc.Numperiods)
}
