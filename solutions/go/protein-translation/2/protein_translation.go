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
        	fallthrough
        case "UUC":
        	return "Phenylalanine", nil
        case "UUA":
        	fallthrough
        case "UUG":
        	return "Leucine", nil
        case "UCU":
        	fallthrough
        case "UCC":
        	fallthrough
        case "UCA":
        	fallthrough
        case "UCG":
        	return "Serine", nil
        case "UAU":
        	fallthrough
        case "UAC":
        	return "Tyrosine", nil
        case "UGU":
        	fallthrough
        case "UGC":
        	return "Cysteine", nil
        case "UGG":
        	return "Tryptophan", nil
        case "UAA":
        	fallthrough
        case "UAG":
        	fallthrough
        case "UGA":
        	return "", ErrStop
        default:
        	return "", ErrInvalidBase
    }
}
