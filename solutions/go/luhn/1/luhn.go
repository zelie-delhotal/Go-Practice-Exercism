package luhn

import (
    "strings"
)

func Valid(a string) bool {
    var id = strings.ReplaceAll(a, " ", "");
    if len(id) <= 1 {
        return false
    }
    var sum int
    var leng = len(id)
    for i:=1; i<=leng; i++ {
        var n = int(id[leng - i]) - int('0')
        if n<0 || n>9 {
            return false
        }
        if i % 2 == 0 {
            n *= 2
            if (n > 9) {
                n -= 9
            }
        }
        sum += n
    }
    return sum % 10 == 0
}
