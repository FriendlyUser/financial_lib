package interest

import (
	"testing"
)

func TestMultiplePeriod(t *testing.T) {
	p := CompoundRateCalculator{Rate: 1}
	value := p.MultiplePeriod(80, 5)
	if value != 2560.0 {
		t.Fatal("Expected value to equal 2560")
	}
}
