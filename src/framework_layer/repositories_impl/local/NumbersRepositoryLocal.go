package local

import (
	"GoChallenge/src/domain_layer/models"
	_ "github.com/go-sql-driver/mysql"
)

type NumbersRepositoryLocal struct {
}

func (this NumbersRepositoryLocal) OrderNumbers(numbers models.Numbers) models.Numbers {
	numbers.Sorted = append(numbers.Sorted)
	arrayNew := make([]int, len(numbers.Unsorted))
	copy(arrayNew, numbers.Unsorted)
	numbers.Sorted = moveItemsRepetToTheEnd(order(arrayNew))
	return numbers
}

func order(input []int) []int {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func moveItemsRepetToTheEnd(input []int) []int {
	n := len(input)
	contChanges := 0
	for i := 0; i < n-1; i++ {
		if (i + contChanges) < n {
			if input[i] == input[i+1] {
				repeated := input[i+1]
				for j := i + 1; j < n-1; j++ {
					input[j] = input[j+1]
				}
				input[n-1] = repeated
				contChanges++
				i--
			}
		}
	}
	return input
}

/*
func order(nums []int) []int {
	a := nums
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}
*/
