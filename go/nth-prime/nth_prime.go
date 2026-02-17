package prime

import (
	"errors"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("there is no zeroth prime")
	}

	if n == 1 {
		return 2, nil
	}

	count := 1
	num := 1

	for count < n {
		num += 2
		if isPrime(num) {
			count++
		}
	}

	return num, nil
}

func isPrime(n int) bool {
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
