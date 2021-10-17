package savings

import (
	"math"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0.0:
		return -3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	default:
		return 0.0 // just to return someting for any case, even if it's unreachable in given confitions
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
	//math.Abs() to get the right sign
	//float64(rate) -- works fine here, but 64 being more precise, the numbers can differ a little
	//in a more complex case it can lead to problems
	return (math.Abs(balance) / 100.0) * float64(rate)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
	return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0 // declared here so it is availiable outside the loop
	for currentBalance := balance; currentBalance <= targetBalance; years++ {
		currentBalance = AnnualBalanceUpdate(currentBalance)
	}
	return years
}
