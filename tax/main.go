package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: tax <base_annual> <employer_pf_annual> <annual_bonus>")
		os.Exit(1)
	}

	base, _ := strconv.ParseFloat(os.Args[1], 64)
	employerPF, _ := strconv.ParseFloat(os.Args[2], 64)
	employeePF := employerPF
	bonus, _ := strconv.ParseFloat(os.Args[3], 64)

	ctc := base + employerPF + bonus

	standardDeduction := 75000.0

	taxableIncome := base + bonus - standardDeduction

	slabs := []struct {
		upperLimit float64
		rate       float64
	}{
		{400000, 0.0},
		{800000, 0.05},
		{1200000, 0.10},
		{1600000, 0.15},
		{2000000, 0.20},
		{2400000, 0.25},
		{9999999999, 0.30},
	}

	incomeTax := 0.0
	previousLimit := 0.0
	for _, slab := range slabs {
		if taxableIncome > slab.upperLimit {
			incomeTax += (slab.upperLimit - previousLimit) * slab.rate
			previousLimit = slab.upperLimit
		} else {
			incomeTax += (taxableIncome - previousLimit) * slab.rate
			break
		}
	}

	cess := incomeTax * 0.04
	totalTax := incomeTax + cess

	// Net = gross earnings - tax - employee PF
	netAnnual := base + bonus - totalTax - employeePF
	monthlyNet := netAnnual / 12

	fmt.Printf("CTC:                    %.2f\n", ctc)
	fmt.Printf("Total tax:              %.2f\n", totalTax)
	fmt.Printf("Annual PF deduction:    %.2f\n", employeePF)
	fmt.Printf("Net annual income:      %.2f\n", netAnnual)
	fmt.Printf("Monthly income (net):   %.2f\n", monthlyNet)
}
