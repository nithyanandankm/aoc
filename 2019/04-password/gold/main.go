package main

import "fmt"

func number_to_array(num int) []int {
	res := make([]int, 0)

	for num != 0 {
		res = append(res, num%10)
		num /= 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func has_incremental_digits(passwd []int) bool {
	for i := 0; i < len(passwd)-1; i++ {
		var sequence = passwd[i+1]
		if passwd[i] > sequence {
			return false
		}
	}

	return true
}

func has_two_equal_adjacent(passwd []int) bool {
	for i := 0; i < len(passwd)-1; i++ {
		if passwd[i] == passwd[i+1] &&
			(i == 0 || passwd[i] != passwd[i-1]) &&
			((i+2) == len(passwd) || passwd[i] != passwd[i+2]) {
			return true
		}
	}

	return false
}

func validate_password_silver(passwd int) bool {
	var passwd_array = number_to_array(passwd)

	return has_incremental_digits(passwd_array) &&
		has_two_equal_adjacent(passwd_array)
}

func main() {
	var lower_limit, upper_limit = 136760, 595730
	matches := 0

	for i := lower_limit; i <= upper_limit; i++ {
		var valid = validate_password_silver(i)

		if valid {
			matches += 1
		}

	}

	fmt.Println("Combinations found: ", matches)
}
