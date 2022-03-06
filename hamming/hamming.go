package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("Unequal length")
    }
	aRune := []rune(a)
    bRune := []rune(b)
	var hammingDistance int
    for i := range aRune {
        if aRune[i] != bRune[i] {
            hammingDistance++
        }
    }
	return hammingDistance, nil
}
