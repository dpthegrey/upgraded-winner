package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dpthegrey/upgraded-winner/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return weight, height
}

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return value
}
