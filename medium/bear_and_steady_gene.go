package medium

import (
	"fmt"
)

const input string = "GAAATAAA"

func getLettersDrop(gene string, geneOfLength int) (map[string]int, int) {
	maxLetter := geneOfLength / 4
	minSize := 0
	result := make(map[string]int)
	result["A"] = 0
	result["C"] = 0
	result["T"] = 0
	result["G"] = 0
	for _, c := range gene {
		letter := fmt.Sprintf("%c", c)
		result[letter] += 1
	}
	for letter, count := range result {
		if count <= maxLetter {
			delete(result, letter)
		} else {
			result[letter] = count - maxLetter
			minSize += result[letter]
		}
	}
	return result, minSize
}

func checkNumberCharInString(s string, c string, number int) bool {
	count := 0
	for _, val := range s {
		char := fmt.Sprintf("%c", val)
		if char == c {
			count += 1
		}
		if count >= number {
			return true
		}
	}
	return false
}

func getMinimumSize(gene string, geneOfLength int, lettersDrop map[string]int, minSize int) int32 {
	for i := 0; i <= geneOfLength-minSize; i++ {
		filterSize := minSize + i
		for pos := 0; pos < geneOfLength-filterSize+1; pos++ {
			subGene := gene[pos : pos+filterSize]
			check := true
			for letter, count := range lettersDrop {
				if checkNumberCharInString(subGene, letter, count) != true {
					check = false
				}
			}
			if check {
				return int32(len(subGene))
			}
		}
	}
	return 0
}

func SteadyGene(gene string) int32 {
	lettersDrop, minSize := getLettersDrop(gene, 8)
	return getMinimumSize(gene, len(gene), lettersDrop, minSize)
}
