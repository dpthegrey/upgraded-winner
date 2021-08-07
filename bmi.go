package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Output information
	fmt.Println("BMI Calculator")
	fmt.Println("--------------------")
	// Prompt for user input (weight + height)
	fmt.Print("Please enter your weight in (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height in (m): ")
	heightInput, _ := reader.ReadString('\n')

	fmt.Print(weightInput)
	fmt.Print(heightInput)
	// Save that user input in variables
	// Calculate BMI (weight / height * height)
	// Output the calculated BMI
}
