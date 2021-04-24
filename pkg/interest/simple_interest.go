package interest

type InterestRateCalculator struct {
	Rate float64
}

// Simple interest formula P = V(1+r).
func (r *InterestRateCalculator) SinglePeriod(value float64) float64 {
	return value * (1 + r.Rate)
}
