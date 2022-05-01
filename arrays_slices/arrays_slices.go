package arrays_slices

func Sum(nums [5]int) (sum int) {
	for _, i := range nums {
		sum += i
	}
	return
}

//using slice
func Sum_v2(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
