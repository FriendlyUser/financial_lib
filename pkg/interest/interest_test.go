package interest_test

import (
	"testing"

	"github.com/FriendlyUser/financial_lib/pkg/interest"
)

func TestSimpleInterest(t *testing.T) {
	p := interest.InterestRateCalculator{Rate: 0.2}
	value := p.SinglePeriod(5)
	if value != 6.0 {
		t.Fatal("Expected value to equal 6")
	}
}
