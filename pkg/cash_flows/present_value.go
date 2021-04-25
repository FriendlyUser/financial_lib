package cash_flows

import "math"

// struct for calculating simple cash flows
type CashFlowCalculator struct {
	Rate         float64
	CashPayments []float64
	TimePeriods  []float64
}

// Formula for present value (PV) for a future payment is determined by
//
// PV = FV / (1 + R) ^ N
//
// PV is the desired present value
//
// FV is the future value that we want to discount
//
// N is the number of periods between present and future value
func (cfc *CashFlowCalculator) CalcPresentValue(futureValue float64, timePeriod float64) float64 {
	pValue := futureValue / math.Pow(1+cfc.Rate, timePeriod)
	return pValue
}

// gets the present value of all cash payments

// see cash_flows.CalcPresentValue
func (cfc *CashFlowCalculator) PresentValue() float64 {
	total := 0.0
	for i := range cfc.CashPayments {
		total += cfc.CalcPresentValue(cfc.CashPayments[i], cfc.TimePeriods[i])
	}
	return total
}

// Adds cash payment
func (cfc *CashFlowCalculator) AddCashPayment(value float64, timePeriod float64) {
	cfc.CashPayments = append(cfc.CashPayments, value)
	cfc.TimePeriods = append(cfc.TimePeriods, timePeriod)
}
