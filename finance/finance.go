// Package finance provides helpers to work with financial needs.
package finance

// SimpleInterest calculates the amount of interest earned over a
// time period using the formula I = Prt.
// p: principal, r: annual interest rate t: period of time
func SimpleInterest(p float64, r float64, t float64) float64 {
	return p * r * t
}

// TotalAmountPlusInterest calculates the total amount plus interest earned
// over a time period.
// p: principal, r: annual interest rate t: period of time
func TotalAmountPlusInterest(p float64, r float64, t float64) float64 {
	interest := SimpleInterest(p, r, t)
	return p + interest
}
