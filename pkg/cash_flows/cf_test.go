package cash_flows_test

import (
	"fmt"
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

func TestAddPaymentsAndPresentValue(t *testing.T) {
	cf := cash_flows.CashFlowCalculator{Rate: 0.25}
	cf.AddCashPayment(1, 1)
	cf.AddCashPayment(2, 2)
	value := cf.PresentValue()
	if value != 2.08 {
		t.Fatal("Expected value to equal 2.08")
	}
}

func TestBondModelling(t *testing.T) {
	bc := cash_flows.BondCalculator{
		Institution: "dli",
		Numperiods:  20,
		Principle:   100000,
		CouponValue: 5000,
	}
	interestRate := bc.InterestRate()
	if interestRate != 0.05 {
		t.Fatal("Expected value to equal 0.5")
	}

	// get formatted string
	niceString := bc.String()
	fmt.Println(niceString)
	if niceString != "dli: at 0.05 for 20" {
		t.Fatal(niceString)
	}
}

// example for cash flow
func ExampleCashFlowCalculator_CalcPresentValue() {
	cf := cash_flows.CashFlowCalculator{Rate: 0.25}
	value := cf.CalcPresentValue(5, 2)
	fmt.Println(value)
}

// example for bond calculator
func ExampleBondCalculator_InterestRate() {
	cf := cash_flows.BondCalculator{
		Institution: "dli",
		Numperiods:  20,
		Principle:   100000,
		CouponValue: 5000,
	}
	value := cf.InterestRate()
	fmt.Println(value)
}
