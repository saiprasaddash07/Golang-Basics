package main

import (
	"fmt"
	"math/rand"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func print(x int) {
	fmt.Println(x)
}

func generateRandom() (int, int) {
	randomnum1 := rand.Intn(20)
	randomnum2 := rand.Intn(20)
	return randomnum1, randomnum2
}

func main() {
	a, b := generateRandom()
	fmt.Println(add(a, b))
}
