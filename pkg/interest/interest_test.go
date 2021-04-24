package interest

import (
	"testing"
)

func TestSimpleInterest(t *testing.T) {
	p := InterestRateCalculator{Rate: 0.2}
	value := p.SinglePeriod(5)
	if value != 6.0 {
		t.Fatal("Expected value to equal 6")
	}
}
