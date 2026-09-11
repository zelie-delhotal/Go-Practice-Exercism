package anagram

import (
    "slices"
    "strings"
    )

func isAnagram(s1 string, s2 string) bool{
    s1 = strings.ToUpper(s1)
    s2 = strings.ToUpper(s2)
    r1 := []rune(s1)
    r2 := []rune(s2)
    slices.Sort(r1)
    slices.Sort(r2)
    return slices.Equal(r1, r2) && s1 != s2
}

func Detect(subject string, candidates []string) []string {
    var res []string
    for _,s := range candidates {
        if isAnagram(subject, s) {
            res = append(res, s)
        }
    }
    return res
}
