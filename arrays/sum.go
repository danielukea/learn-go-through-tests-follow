package main

func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum = sum + number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
