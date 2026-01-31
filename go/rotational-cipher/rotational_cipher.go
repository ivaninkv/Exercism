package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = ((shiftKey % 26) + 26) % 26

	bytes := []byte(plain)
	for i, b := range bytes {
		switch {
		case b >= 'a' && b <= 'z':
			bytes[i] = 'a' + (b-'a'+byte(shiftKey))%26
		case b >= 'A' && b <= 'Z':
			bytes[i] = 'A' + (b-'A'+byte(shiftKey))%26
		}
	}
	return string(bytes)
}
