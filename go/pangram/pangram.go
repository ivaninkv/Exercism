package pangram

// func IsPangram(input string) bool {
// 	alphabet := [26]int{}

// 	for _, r := range input {
// 		if r >= 'a' && r <= 'z' {
// 			alphabet[r-'a']++
// 		}
// 		if r >= 'A' && r <= 'Z' {
// 			alphabet[r+32-'a']++
// 		}
// 	}

// 	for _, a := range alphabet {
// 		if a == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsPangram(input string) bool {
	var mask uint32

	for _, r := range input {
		lower := r | 32
		if lower >= 'a' && lower <= 'z' {
			mask |= 1 << (lower - 'a')
			if mask == 0x3FFFFFF {
				return true
			}
		}
	}

	return mask == 0x3FFFFFF
}
