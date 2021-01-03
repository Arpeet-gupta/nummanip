package calc

import "errors"

// return sum of two integers with error
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum
	}
}

func Multi(numbers ...int) (error, int) {
	multi := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 numbers"), multi
	} else {
		for _, num := range numbers {
			multi = multi * num
		}
		return nil, multi
	}
}
func Div(num1 int, num2 int) int {
	div := 0
	div = num1 / num2
	return div
}
