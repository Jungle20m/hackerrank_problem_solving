package main

import "fmt"

func formingMagicSquare(square [][]int32) int32 {
	squareSize := len(square)

	for row := 0; row < squareSize; row++ {
		for column := 0; column < squareSize; column++ {
			println(square[row][column])
		}
	}

	return 0
}

func main() {
	square := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	cost := formingMagicSquare(square)
	fmt.Println(cost)
}
