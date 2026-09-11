package sumofmultiples

import "slices"

func SumMultiples(limit int, divisors ...int) int {
    var multiples []int
    for _,n := range divisors {
        if n == 0 {
            continue
        }
        var m = n
        for m < limit {
            multiples = append(multiples, m)
            m += n
        }
    }
    slices.Sort(multiples)
    multiples = slices.Compact(multiples)
    var ret int
    for _,n := range multiples {
        ret += n
    }
    return ret
}
