package medium

func formingMagicSquare(square [][]int32) int32 {
	squareSize := len(square)

	for row := 0; row < squareSize; row++ {
		for column := 0; column < squareSize; column++ {
			println(square[row][column])
		}
	}

	return 0
}
