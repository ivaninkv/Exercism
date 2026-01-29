package isbn

func IsValidISBN(isbn string) bool {
	var sum, digits int
	for _, c := range isbn {
		if c == '-' {
			continue
		}
		if digits > 9 { // More than 10 digits (or 'X') found
			return false
		}
		var d int
		if c >= '0' && c <= '9' {
			d = int(c - '0')
		} else if c == 'X' && digits == 9 {
			d = 10
		} else {
			return false // Invalid character
		}
		sum += d * (10 - digits)
		digits++
	}
	return digits == 10 && sum%11 == 0
}
