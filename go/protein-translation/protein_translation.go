package protein

import (
	"errors"
)

var (
	ErrStop = errors.New("stop")

	ErrInvalidBase = errors.New("invalid codon")

	noProteins = []string{}
)

const codonLength = 3

func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return protein, err
}

func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += codonLength {
		protein, err := FromCodon(rna[i : i+codonLength])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return noProteins, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, err
}
