package interest

// struct for calculating compound rates
type MACalculator struct {
	MNumPeriod int
	mPrices    []float64
}

func (ma *MACalculator) CalculateMA() []float64 {
	var ma_values []float64
	sum := 0.0
	for i := range ma.mPrices {
		sum += ma.mPrices[i]
		if i > ma.MNumPeriod {
			temp := sum / float64(ma.MNumPeriod)
			ma_values = append(ma_values, temp)
		}
	}
	return ma_values
}

func (ma *MACalculator) CalculateEMA() []float64 {
	var ema_values []float64
	sum := 0.0
	multipler := 2.0 / float64(ma.MNumPeriod+1)
	for i := range ma.mPrices {
		sum += ma.mPrices[i]
		if i == ma.MNumPeriod {
			ema_values = append(ema_values, sum/float64(ma.MNumPeriod))
			sum -= ma.mPrices[i-ma.MNumPeriod]
		} else if i > ma.MNumPeriod {
			last_entry := ema_values[len(ema_values)-1]
			val := (1-multipler)*last_entry + multipler*ma.mPrices[i]
			ema_values = append(ema_values, val)
		}
	}
	return ema_values
}

func (ma *MACalculator) AddPriceQuote(close float64) {
	ma.mPrices = append(ma.mPrices, close)
}
