package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const separator = "----------------------------------------"
const weightPrompt = "Please enter your weight in (kg): "
const heightPrompt = "Please enter your height in (m): "

func main() {
	// Output information
	fmt.Println(mainTitle)
	fmt.Println(separator)
	// Prompt for user input (weight + height)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	fmt.Print(weightInput)
	fmt.Print(heightInput)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculate BMI (weight / height * height)
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is: %.2f", bmi)
}
