package main

import (
	"errors"
	"fmt"
	"os"
)

const resultFile = "result.txt"

func main() {

	revenue, err1 := userInput("Inserte Revenue: ")

	expenses, err2 := userInput("Inserte Expenses: ")

	taxRate, err3 := userInput("Insert Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("error")
		fmt.Println(err1)
		fmt.Println("--")
		return
	}

	etb, profit, ratio := userCalculation(revenue, expenses, taxRate)

	fmt.Printf("ETB= %.1f Profit= %.1f  Ratio= %.1f  ", etb, profit, ratio)
	writeResultToFile(etb, profit, ratio)
}

func userInput(text string) (float64, error) {

	var input float64
	fmt.Print(text)
	fmt.Scan(&input)
	if input <= 0 {

		return 0, errors.New("el valor no puede ser menor o igual a cero ")

	}

	return input, nil

}

func userCalculation(revenue, expenses, taxRate float64) (etb float64, profit float64, ratio float64) {
	etb = revenue - expenses
	profit = etb * (1 - taxRate/100)
	ratio = etb / profit
	return
}

func writeResultToFile(etb, profit, ratio float64) {
	resultText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f", etb, profit, ratio)
	os.WriteFile(resultFile, []byte(resultText), 0644)

}
