package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(rand.Intn(100) + 1)
	/* using rand.Intn(100) + 1 is a way to generate a random integer between 1 and 100, whereas using rand.Intn(101) generates a random integer between 0 and 100. _chatGPT*/
	//seconds := time.Now().Unix()
	//fmt.Println(time.Now())
	//fmt.Println(seconds)
	//rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		if guesses == 10 {
			fmt.Println("You lost the game. It was", target)
			break
		}
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Your guess was HIGH.")
		} else if guess == target {
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
}
