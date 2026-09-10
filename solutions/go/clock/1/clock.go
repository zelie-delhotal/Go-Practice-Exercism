package clock

import (
    "fmt"
)

type Clock struct {
    h int
    m int
}

func New(h, m int) Clock {
    var c Clock
    c.h = h
    c.m = m
    return c.normalize()
}

func (c Clock) Add(m int) Clock {
    c.m += m
    return c.normalize()
}

func (c Clock) Subtract(m int) Clock {
    c.m -= m
    return c.normalize()
}

func (c Clock) normalize() Clock{
    if c.m < 60 && c.m >= 0 && c.h < 24 && c.h >= 0{
        return c
    }
    if c.m >= 60 {
        c.m -= 60
        c.h ++
    }
    if (c.m < 0) {
        c.m += 60
        c.h --
    }
    if (c.h >= 24) {
        c.h -= 24
    }
    if (c.h < 0) {
        c.h += 24
    }
    return c.normalize()
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
