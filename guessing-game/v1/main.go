package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 1001
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
