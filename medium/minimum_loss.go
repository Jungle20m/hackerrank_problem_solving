package medium

import (
	"sort"
)

// Input:
// [20 7 8 2 5]

type Pair struct {
	Key   int
	Value int
}

type Pairs []Pair

func MinimumLoss(price []int64) int32 {
	var min int = 2*10e5 + 1

	pairs := make(Pairs, 0, len(price))

	for k, v := range price {
		//m[k] = int(v)
		pairs = append(pairs, Pair{k, int(v)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	for i := 0; i < len(pairs)-1; i++ {
		if pairs[i+1].Value-pairs[i].Value < min && pairs[i+1].Key < pairs[i].Key {
			min = pairs[i+1].Value - pairs[i].Value
		}
	}

	return int32(min)
}
