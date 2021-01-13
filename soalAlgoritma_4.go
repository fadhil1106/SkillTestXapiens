package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	var number = rand.Intn(max - min + 1)

	fmt.Println("Random Number : ", number)
	guess(number)
}

func guess(number int) {
	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(100)
	count := 1
	maxGuess := 100
	minGuess := 0
	fmt.Printf("min:%d max:%d guess:%d \n", minGuess, maxGuess, guess)

	for guess != number {
		if guess > number {
			if guess < maxGuess {
				maxGuess = guess
			}
		} else {
			if guess > minGuess {
				minGuess = guess
			}
		}
		guess = int(math.Ceil(float64(maxGuess-minGuess)/2) + float64(minGuess))
		count++
		fmt.Printf("min:%d max:%d guess:%d \n", minGuess, maxGuess, guess)
	}
	fmt.Print("Jumlah Langkah : ", count)
}
