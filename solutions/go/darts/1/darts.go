package darts

import (
    "math"
)

func Score(x, y float64) int {
    var distance float64 = math.Sqrt(x*x + y*y)
    if distance <= 1.0 {
        return 10
    }
    if distance <= 5.0 {
        return 5
    }
    if distance <= 10.0 {
        return 1
    }
    return 0
}
