package arraySlices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sumList []int) {
	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		sumList = append(sumList, sum)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sumList []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumList = append(sumList, 0)
		} else {
			tail := numbers[1:]
			sumList = append(sumList, Sum(tail))
		}
	}
	return
}
