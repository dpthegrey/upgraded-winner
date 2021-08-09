package main

import (
	"github.com/dpthegrey/upgraded-winner/info"
)

func main() {
	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}
