package easy

func getNextMultipleOfFive(number int32) int32 {
	remainder := number % 5
	return number + 5 - remainder
}

func GradingStudents(grades []int32) []int32 {
	var result []int32
	for _, grade := range grades {
		nextMultipleOf5 := getNextMultipleOfFive(grade)
		if grade < 38 {
			result = append(result, grade)
			continue
		}
		if nextMultipleOf5-grade < 3 {
			result = append(result, nextMultipleOf5)
			continue
		}
		result = append(result, grade)
	}
	return result
}
