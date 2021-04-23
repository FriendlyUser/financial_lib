package cash_flows

import "math"

type CashFlowCalculator struct {
	rate float64
	cashPayments []float64
	timePeriods []float64
}

// Formula for present value (PV) for a future payment is determined by
// PV = FV / (1 + R) ^ N
// PV is the desired present value
// FV is the future value that we want to discount
// N is the number of periods between present and future value
func (cfc *CashFlowCalculator) CalcPresentValue(futureValue float64, timePeriod float64) float64 {
	pValue := futureValue / math.Pow(1 + cfc.rate, timePeriod)
	return pValue
}

// gets the present value of all cash payments
// see cash_flows.CalcPresentValue
func (cfc *CashFlowCalculator) PresentValue() float64 {
	total := 0.0
	for i, _ := range cfc.cashPayments {
		total += cfc.CalcPresentValue(cfc.cashPayments[i], cfc.timePeriods[i])
	}
	return total
}

// Adds cash payment 
func (cfc *CashFlowCalculator) AddCashPayment(value float64, timePeriod float64) {
	cfc.cashPayments = append(cfc.cashPayments, value)
	cfc.timePeriods = append(cfc.timePeriods, timePeriod)
}
