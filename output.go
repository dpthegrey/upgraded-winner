package main

import "fmt"

func printBMI(bmi float64) {
	if bmi < 18.5 {
		fmt.Printf("your BMI is %.2f. You are underweight", bmi)
	} else if bmi < 25 {
		fmt.Printf("your BMI is %.2f. You are normal", bmi)
	} else if bmi < 30 {
		fmt.Printf("your BMI is %.2f. You are overweight", bmi)
	} else {
		fmt.Printf("your BMI is %.2f. You are obese", bmi)
	}
}
