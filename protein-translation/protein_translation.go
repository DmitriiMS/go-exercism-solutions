package protein

import (
	"errors"
)

var err error = nil
var ErrStop error = errors.New("stop")
var ErrInvalidBase error = errors.New("u wot m8?")

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", err
	case "UUU", "UUC":
		return "Phenylalanine", err
	case "UUA", "UUG":
		return "Leucine", err
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", err
	case "UAU", "UAC":
		return "Tyrosine", err
	case "UGU", "UGC":
		return "Cysteine", err
	case "UGG":
		return "Tryptophan", err
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	result := []string{}
LOOP:
	for start, end := 0, 3; end <= len(rna); start, end = start+3, end+3 {
		protein, codonErr := FromCodon(rna[start:end])
		switch codonErr {
		case ErrInvalidBase:
			return result, codonErr
		case ErrStop:
			break LOOP
		}
		result = append(result, protein)
	}
	return result, nil
}
