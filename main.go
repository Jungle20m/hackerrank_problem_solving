package main

import (
	"fmt"
	"problem_solving/medium"
)

func main() {
	price := []int64{20, 7, 8, 2, 5}
	result := medium.MinimumLoss(price)
	fmt.Println(result)
}
