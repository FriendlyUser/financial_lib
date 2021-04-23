package interest

import "math"

type CompoundRateCalculator struct {
	rate float64
}

// Formula for discrete	compound interest is
// V = P(1+R)^N
// where P is the present value
// R is the rate
// N is the number of periods
func (r *CompoundRateCalculator) MultiplePeriod(value float64, numPeriods float64) float64 {
	return value * math.Pow(1 + r.rate, numPeriods)
}


// The formula for the calculation of continuous interest rate compounding is
// V = Pe^(RN)
// V is desired future value
// P is the present value
// R is the interest rate during this period
// N is the number of periods
func (r *CompoundRateCalculator) ContinuousCompounding(value float64, numPeriods float64) float64 {
	return value * math.Exp(r.rate * numPeriods)
}