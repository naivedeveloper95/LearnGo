package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const promtp = " and don't type your number in, just press ENTER."

func main() {
	// Seed the random number generator

	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer int

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

	answer = (firstNumber * secondNumber) - substraction
	fmt.Println(answer)

}
