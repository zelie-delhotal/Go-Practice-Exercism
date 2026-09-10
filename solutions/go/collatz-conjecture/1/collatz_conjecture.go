// Package collatzconjecture contains tools regarding the Collatz conjecture.
package collatzconjecture

import (
    "errors"
)
// Function CollatzConjecture returns the number of iterations to get from a number
// given as parameter to 1 following the rules of the Collatz conjecture.
func CollatzConjecture(n int) (int, error) {
    var tempN int
    var tempE error
    if (n <= 0) {
        return 0, errors.New("n must be a positive integer.")
    }
	if n != 1 {
        if n % 2 == 0 {
            tempN, tempE = CollatzConjecture(n / 2)
        } else {
            tempN, tempE = CollatzConjecture(3*n + 1)
        }
        return 1+tempN, tempE;
    }
    return 0, nil
}
