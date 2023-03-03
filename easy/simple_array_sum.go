package easy

func SimpleArraySum(arr []int32) int32 {
	var result int32 = 0
	for _, ele := range arr {
		result += ele
	}
	return result
}
