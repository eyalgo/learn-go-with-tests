package main

func _sum(numbers []int, tailIndex int) int {
	tail := numbers[tailIndex:]
	sum := 0
	for _, number := range tail {
		sum += number
	}
	return sum
}

func Sum(numbers []int) int {
	return _sum(numbers, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, _sum(numbers, 0))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, _sum(numbers, 1))
		}
	}
	return sums
}
