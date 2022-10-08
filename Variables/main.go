package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = " and don't type your number in, just press ENTER."

func main() {
	// Seed the random number generator

	// rand generates a number between 0 and whatever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and go
	// greater than 10 (multiplying by 1 makes the game a bit silly)
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = (firstNumber * secondNumber) - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)

}

func playTheGame(firstNumber int, secondNumber int, substraction int, answer int) {
	// display the wecome instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")

	fmt.Println("")

	// create a reader variable which reads the input from the console
	reader := bufio.NewReader(os.Stdin)

	// Take the player through the steps of the game
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract ", substraction, prompt)
	reader.ReadString('\n')

	// Print the answer on the console.
	fmt.Println(answer)
}
