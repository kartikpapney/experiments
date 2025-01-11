package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	currentCtcInput := os.Args[1]
	ctcInput := os.Args[2]
	ctc, _ := strconv.ParseFloat(ctcInput, 64)
	currentCtc, _ := strconv.ParseFloat(currentCtcInput, 64)

	standardDeduction := 75000.0
	pf := 1800.0*12

	slabs := []struct {
		upperLimit float64
		rate       float64
	}{
		{300000, 0.0},      // No tax for income up to 2.5 lakh
		{700000, 0.05},     // 5% tax for income between 3 lakh to 7 lakh
		{1000000, 0.10},    // 10% tax for income between 7 lakh to 10 lakh
		{1200000, 0.15},    // 15% tax for income between 10 lakh to 12 lakh
		{1500000, 0.20},    // 20% tax for income between 12 lakh to 15 lakh
		{9999999999, 0.30}, // 30% tax for income above 15 lakh
	}

	taxableIncome := ctc - standardDeduction - pf

	totalTax := 0.0
	previousLimit := 0.0

	for _, slab := range slabs {
		if taxableIncome > slab.upperLimit {
			totalTax += (slab.upperLimit - previousLimit) * slab.rate
			previousLimit = slab.upperLimit
		} else {
			totalTax += (taxableIncome - previousLimit) * slab.rate
			break
		}
	}

	netAnnualIncome := ctc - totalTax - pf*2
	monthlyIncome := netAnnualIncome / 12

	fmt.Printf("Tax you will pay: %.2f\n", totalTax);
	fmt.Printf("Monthly Income after tax: %.2f\n", monthlyIncome)
	fmt.Printf("Increment percent: %.2f\n", 100.0*(ctc-currentCtc)/currentCtc)
}
