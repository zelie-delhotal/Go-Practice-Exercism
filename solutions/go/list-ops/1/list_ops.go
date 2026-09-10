// Package listops creates a IntList type with corresponding list operations
package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
    for _,i := range s {
        initial = fn(initial, i)
    }
    return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    s = s.Reverse()
    for _,i := range s {
        initial = fn(i, initial)
    }
    return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
    var res IntList
    for _,i := range s {
        if fn(i) {
            res = append(res, i)
        }
    }
    return res
}

func (s IntList) Length() int {
    return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
    for index,i := range s {
        s[index] = fn(i)
    }
    return s
}

func (s IntList) Reverse() IntList {
    var res IntList
    for _,i := range s {
        res = append([]int{i}, res...)
    }
    return res
}

func (s IntList) Append(lst IntList) IntList {
    s = append(s, lst...)
    return s
}

func (s IntList) Concat(lists []IntList) IntList {
    var res IntList
    for _,l := range lists {
        res = append (res, l...)
    }
    return res
}
