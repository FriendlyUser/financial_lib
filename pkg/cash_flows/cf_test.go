package cash_flows_test

import (
	"testing"

	"github.com/FriendlyUser/financial_lib/pkg/cash_flows"
)

func TestCashFlowCalculator(t *testing.T) {
	cf := cash_flows.CashFlowCalculator{Rate: 0.25}
	value := cf.CalcPresentValue(5, 2)
	if value != 3.2 {
		t.Fatal("Expected value to equal 3.2")
	}
}
