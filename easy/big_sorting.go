package easy

import "fmt"

func compare(str1, str2 string) int {
	if len(str1) < len(str2) {
		return -1
	}
	if len(str1) > len(str2) {
		return 1
	}
	if str1 > str2 {
		return 1
	} else if str1 < str2 {
		return -1
	} else {
		return 0
	}
}

func BigSorting(unsorted []string) []string {
	var result []string

	for _, element := range unsorted {
		fmt.Println(element)
	}

	return result
}
