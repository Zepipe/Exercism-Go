package protein

import (
	"errors"
)

// Define the expected error variables
var (
	ErrStop        = errors.New("stop codon encountered")
	ErrInvalidBase = errors.New("invalid codon")
)

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	
	// Process RNA in chunks of 3 characters (codons)
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			return nil, ErrInvalidBase
		}
		
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		
		if err != nil {
			if err == ErrStop {
				// STOP codon means we're done, return what we have
				return proteins, nil
			}
			return nil, err
		}
		
		proteins = append(proteins, protein)
	}
	
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonToProtein[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}