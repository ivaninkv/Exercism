package strand

func ToRNA(dna string) string {
	var output = make([]byte, len(dna))
	for i, r := range dna {
		switch r {
		case 'G':
			output[i] = 'C'
		case 'C':
			output[i] = 'G'
		case 'T':
			output[i] = 'A'
		case 'A':
			output[i] = 'U'
		default:
			panic("Invalid DNA nucleotide")
		}
	}
	return string(output)
}
