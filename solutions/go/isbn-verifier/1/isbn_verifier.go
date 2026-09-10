package isbnverifier


func IsValidISBN(isbn string) bool {
    var hyphens, sum int
    for i:= range isbn {
        if (isbn[i] >= '0' && isbn[i] <= '9') || isbn[i] == 'X' {
            sum++
        }
    }
    if sum != 10 {
        return false
    }
    sum = 0
    for i,c:= range isbn {
        if (c == '-'){
            if (i - hyphens == 0) {
                return false
            } else {
                hyphens++
                continue
            }
        }
        if c < '0' || c > '9' {
            if i-hyphens == 9 && c == 'X'{
                sum += 10
            } else {
                return false
            }
        } else if i - hyphens < 10 {
            sum += (int(c) - '0') * (10 - i + hyphens)
        } else {
            return false
        }
    }
    return sum % 11 == 0
}
