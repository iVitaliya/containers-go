package utils

import "fmt"

func DefaultEquals[T any](a T, b T) bool {
	A := fmt.Sprint(a)
	B := fmt.Sprint(b)

	return A == B
}
