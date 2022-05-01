package arrays_slices

//using array
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

func Sum_v1(nums ...[]int) []int {
	var sum []int = make([]int, len(nums))
	for i, n := range nums {
		sum[i] = Sum_v2(n)
	}
	return sum
}

func sumTails(nums ...[]int) (sum []int) {
	for _, arr := range nums {
		if len(arr) > 1 {
			tail := arr[1:]
			sum = append(sum, Sum_v2(tail))
		} else {
			sum = append(sum, 0)
		}

	}
	return
}
