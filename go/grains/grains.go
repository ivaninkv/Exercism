package grains

import "fmt"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("invalid square number")
	}
	return uint64(1 << (number - 1)), nil
}

func Total() uint64 {
	total := uint64(0)
	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		total += grains
	}
	return total
	// return 1<<64 - 1 // the best way to calculate the total grains on a chessboard
}
