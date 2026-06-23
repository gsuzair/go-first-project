package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmout float64 = 1000
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmout)

	outputText("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	calculatedFV, calculatedFRV := calculateFutureValue(investmentAmout, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", calculatedFV)
	formattedFRV := fmt.Sprintf("Future Real Value: %.1f\n", calculatedFRV)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmout, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmout * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
