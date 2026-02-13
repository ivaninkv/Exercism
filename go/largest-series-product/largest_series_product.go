package lsproduct

import "fmt"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	err := checkInput(digits, span)
	if err != nil {
		return 0, err
	}

	maxProduct := int64(0)
	for i := 0; i <= len(digits)-span; i++ {
		product := int64(1)
		for j := i; j < i+span; j++ {
			digit := int64(digits[j] - '0')
			product *= digit
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}

func checkInput(digits string, span int) error {
	if span > len(digits) {
		return fmt.Errorf("span must be smaller than string length")
	}
	if span < 0 {
		return fmt.Errorf("span must not be negative")
	}

	for _, ch := range digits {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("digits input must only contain digits")
		}
	}

	return nil
}
