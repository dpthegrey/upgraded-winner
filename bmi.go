package main

import (
	"fmt"

	"github.com/dpthegrey/upgraded-winner/info"
)

func main() {
	// Output information
	info.PrintWelcome()

	weight, height := getUserMetrics()

	// Calculate BMI (weight / height * height)
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is: %.2f", bmi)
}
