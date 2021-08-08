package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	// Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	fmt.Print(weightInput)
	fmt.Print(heightInput)

	weight, err := strconv.ParseFloat(weightInput, 64)
	height, err := strconv.ParseFloat(heightInput, 64)

	// Calculate BMI (weight / height * height)
	// Output the calculated BMI
}
