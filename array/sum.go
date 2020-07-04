package array

func Sum(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	result := []int{}
	for _, number := range numbersToSum {
		result = append(result, Sum(number))
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	result := []int{}
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			result = append(result, Sum(numbers[1:]))
		} else {
			result = append(result, 0)
		}
	}
	return result
}
