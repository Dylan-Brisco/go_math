package geometry

import "errors"

func Cube(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("zero length edge is not allowed")
}
