package equities_test

import (
	"fmt"
	"testing"

	"github.com/FriendlyUser/financial_lib/pkg/equities"
)

func TestExampleMACalculator_CalculateMA(t *testing.T) {
	ma := equities.MACalculator{MNumPeriod: 2}
	ma.AddPriceQuote(2)
	ma.AddPriceQuote(5)
	ma.AddPriceQuote(7)
	ma.AddPriceQuote(9)
	values := ma.CalculateMA()
	fmt.Printf("%v\n", values)
	if values[0] != 11.5 {
		t.Errorf("Expecting 11.5")
	}
}

func TestExampleMACalculator_CalculateEMA(t *testing.T) {
	ma := equities.MACalculator{MNumPeriod: 2}
	ma.AddPriceQuote(2)
	ma.AddPriceQuote(5)
	ma.AddPriceQuote(7)
	ma.AddPriceQuote(9)
	values := ma.CalculateEMA()
	fmt.Printf("%v\n", values)
	if values[0] != 7 {
		t.Errorf("Expecting 11.5")
	}
}

// Examples
func ExampleMACalculator_CalculateMA() {
	ma := equities.MACalculator{MNumPeriod: 2}
	ma.AddPriceQuote(2)
	ma.AddPriceQuote(5)
	ma.AddPriceQuote(7)
	ma.AddPriceQuote(9)
	values := ma.CalculateMA()
	fmt.Printf("%v\n", values)
}
