package array

func Sum(numbers []int) int {
	sum := 0
	if len(numbers) == 0 {
		return sum
	}
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	tempToSum := [][]int{}
	for _, v := range numbersToSum {
		if len(v) == 0 {
			tempToSum = append(tempToSum, []int{0})
		} else {
			tempToSum = append(tempToSum, v[1:])
		}

	}
	return SumAll(tempToSum...)
}
