package interest

type InterestRateCalculator struct {
	rate float64
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *InterestRateCalculator) SinglePeriod(value float64) float64 {
	return value * (1 + r.rate)
}


// func main() {
// 	r := rect{width: 10, height: 5}

// 	// Here we call the 2 methods defined for our struct.
// 	fmt.Println("area: ", r.area())
// 	fmt.Println("perim:", r.perim())

// 	// Go automatically handles conversion between values
// 	// and pointers for method calls. You may want to use
// 	// a pointer receiver type to avoid copying on method
// 	// calls or to allow the method to mutate the
// 	// receiving struct.
// 	rp := &r
// 	fmt.Println("area: ", rp.area())
// 	fmt.Println("perim:", rp.perim())
// }

