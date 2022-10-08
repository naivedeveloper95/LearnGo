package main

import (
	"bufio"
	"fmt"
	"os"
)

const promtp = " and press enter when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var substraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")

	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", promtp)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, promtp)
	reader.ReadString('\n')

	fmt.Println("Multiply the result by", secondNumber, promtp)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", promtp)
	reader.ReadString('\n')

	fmt.Println("Now subtract ", substraction, promtp)
	reader.ReadString('\n')

}
