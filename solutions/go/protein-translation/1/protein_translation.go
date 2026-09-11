package proteintranslation

import (
    "errors"
)

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
    var ret []string
    var runes = []rune(rna)
    for i,_ := range runes {
        if i%3 == 0 {
            if i+3 > len(rna) {
                return nil, ErrInvalidBase
            }
        	s, err := FromCodon(string(runes[i:i+3]))
            if err == nil {
                ret = append(ret, s)
            } else if err == ErrStop {
                break
            } else {
                return nil, err
            }
        }
    }
    if len(ret) == 0 {
        return nil, nil
    }
    return ret, nil
}

func FromCodon(codon string) (string, error) {
    switch codon {
        case "AUG":
        	return "Methionine", nil
        case "UUU":
        	return "Phenylalanine", nil
        case "UUC":
        	return "Phenylalanine", nil
        case "UUA":
        	return "Leucine", nil
        case "UUG":
        	return "Leucine", nil
        case "UCU":
        	return "Serine", nil
        case "UCC":
        	return "Serine", nil
        case "UCA":
        	return "Serine", nil
        case "UCG":
        	return "Serine", nil
        case "UAU":
        	return "Tyrosine", nil
        case "UAC":
        	return "Tyrosine", nil
        case "UGU":
        	return "Cysteine", nil
        case "UGC":
        	return "Cysteine", nil
        case "UGG":
        	return "Tryptophan", nil
        case "UAA":
        	return "", ErrStop
        case "UAG":
        	return "", ErrStop
        case "UGA":
        	return "", ErrStop
        default:
        	return "", ErrInvalidBase
    }
}
