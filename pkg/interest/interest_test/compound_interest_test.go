package interest_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/FriendlyUser/financial_lib/pkg/interest"
)

func TestMultiplePeriod(t *testing.T) {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.MultiplePeriod(2, 5)
	if value != 64 {
		t.Fatal("Expected value to equal 64")
	}
}

func TestContinuousCompounding(t *testing.T) {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.ContinuousCompounding(0.5, 2)
	if math.Round(value) != 11 {
		t.Fatal("Expected value to equal 2560")
	}
}

// Example code
func ExampleCompoundRateCalculator_ContinuousCompounding() {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.ContinuousCompounding(0.5, 2)
	fmt.Printf("Result is %f", value)
}

func ExampleCompoundRateCalculator_MultiplePeriod() {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.MultiplePeriod(0.5, 2)
	fmt.Printf("Result is %f", value)
}
