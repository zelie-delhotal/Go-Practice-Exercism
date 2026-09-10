// Package collatzconjecture contains tools regarding the Collatz conjecture.
package collatzconjecture

import "errors"

// Function CollatzConjecture returns the number of iterations to get from a number
// given as parameter to 1 following the rules of the Collatz conjecture.
func CollatzConjecture(n int) (int, error) {
    var nbIterations int
    if n <= 0 {
        return 0, errors.New("n must be a positive integer.")
    }
    for n > 1 {
        nbIterations ++
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n * 3 + 1
        }
    }
    return nbIterations, nil
}
