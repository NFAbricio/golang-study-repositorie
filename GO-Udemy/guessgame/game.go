package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(
		"A random number wiil be sorted between 1 and 100. Try to guess it!",
	)

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guess := [10]int64{}

	for i := range guess {
		fmt.Println("Whats your guess?")
		scanner.Scan()
		shot := scanner.Text()
		shot = strings.TrimSpace(shot)

		guessInt, err := strconv.ParseInt(shot, 10, 64)
		if err != nil {
			fmt.Println("Your try must be a integer number")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("Your guess is lower than the number - your guess ->", guessInt)
		case guessInt > x:
			fmt.Println("Your guess is higher than the number - your guess ->", guessInt)
		case guessInt == x:
		fmt.Printf("\nCongrats! You won! The number was %d\n"+
		"You took %d tries to guess it!\n"+
		"You has 10 chances and you guess was %v\n", x, i+1, guess[:i])
			return
		}

		guess[i] = guessInt
	}

}