package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spdash/bmi/info"
)

func main() {
	// Output information
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	// Prompt for user input (weight+height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	// fmt.Print(weightInput, heightInput)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculte the BMI i.e (weight / (height * height))
	bmi := weight / (height * height)

	// Ouput the calculated bmi
	fmt.Printf("Your bmi: %.2f", bmi)
}
