package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    var out = make(map[string]int)
    for i,list := range in {
        for _,s := range list {
            out[strings.ToLower(s)] = i
        }
    }
    return out
}
