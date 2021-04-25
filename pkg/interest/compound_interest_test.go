package interest_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/FriendlyUser/financial_lib/pkg/interest"
)

func TestMultiplePeriod(t *testing.T) {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.MultiplePeriod(80, 5)
	if value != 2560.0 {
		t.Fatal("Expected value to equal 2560")
	}
}

func TestContinuousCompounding(t *testing.T) {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.ContinuousCompounding(0.5, 2)
	if math.Round(value) != 4 {
		t.Fatal("Expected value to equal 4")
	}
}

// Example code
func ExampleCompoundRateCalculator_ContinuousCompounding() {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.ContinuousCompounding(0.5, 2)
	fmt.Printf("Result is %f", value)
}

// example code please do something
func ExampleCompoundRateCalculator_MultiplePeriod() {
	p := interest.CompoundRateCalculator{Rate: 1}
	value := p.MultiplePeriod(0.5, 2)
	fmt.Printf("Result is %f", value)
}
