package medium

import (
	"fmt"
)

// Example input
// ("GAAATAAA", 8)
// ("TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC", 40)

func SteadyGene(gene string, length int32) int32 {
	m := make(map[string]int)
	m["A"] = 0
	m["C"] = 0
	m["G"] = 0
	m["T"] = 0

	for _, c := range gene {
		letter := fmt.Sprintf("%c", c)
		m[letter] += 1
	}

	var left int32
	var right int32
	min := length
	steady := len(gene) / 4

	for true {
		for m["A"] > steady || m["C"] > steady || m["G"] > steady || m["T"] > steady {
			letter := fmt.Sprintf("%c", gene[right])
			m[letter]--
			right++

			if right == length {
				break
			}
		}
		if right == length {
			break
		}

		for m["A"] <= steady && m["C"] <= steady && m["G"] <= steady && m["T"] <= steady {
			letter := fmt.Sprintf("%c", gene[left])
			m[letter]++
			left++
		}

		if right-left+1 < min {
			min = right - left + 1
		}
	}

	return min
}
