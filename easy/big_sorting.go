package easy

import (
	"sort"
)

func BigSorting(unsorted []string) []string {
	sort.Slice(unsorted, func(i, j int) bool {
		if len(unsorted[i]) != len(unsorted[j]) {
			return len(unsorted[i]) < len(unsorted[j])
		}
		return unsorted[i] < unsorted[j]
	})
	return unsorted
}
