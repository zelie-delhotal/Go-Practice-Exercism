package phonenumber

import (
    "errors";
    "strings";
    "fmt"
)

func Number(s string) (string, error) {
    s = strings.Replace(s, ".", "", 2)
    s = strings.Replace(s, "-", "", 2)
    s = strings.Replace(s, "(", "", 1)
    s = strings.Replace(s, ")", "", 1)
    s = strings.ReplaceAll(s, " ", "")
    if len(s) == 11 {
    	s = strings.TrimPrefix(s, "1")
    }
    if len(s) == 12 {
    	s = strings.TrimPrefix(s, "+1")
    }
    r := []rune(s)
    for i, c := range r {
        if (c == '1' || c == '0')&& (i == 0 || i == 3) {
            return "", errors.New("")
        }
        if (c < '0' || c > '9' || i > 9){
            return "", errors.New("")
        }
    }
    return s, nil
}

func AreaCode(s string) (string, error) {
    var cleaned string
    var e error
    cleaned, e = Number(s)
    if e != nil {
        return "", e
    }
    r := []rune(cleaned)
    return string(r[:3]), nil
}

func Format(s string) (string, error) {
    var cleaned string
    var e error
    cleaned, e = Number(s)
    if e != nil {
        return "", e
    }
    r := []rune(cleaned)
    return fmt.Sprintf("(%s) %s-%s", string(r[:3]), string(r[3:6]), string(r[6:10])), nil
}
