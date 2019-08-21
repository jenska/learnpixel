package core

import "fmt"

func Zip(a, b []interface{}) ([]interface{}, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}
	result := make([]interface{}, len(a)+len(b))
	for i := range a {
		result[i*2] = a[i]
		result[i*2+1] = b[i]
	}
	return result, nil
}

func IntZip(a, b []int) ([]int, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}
	result := make([]int, len(a)+len(b))
	for i := range a {
		result[i*2] = a[i]
		result[i*2+1] = b[i]
	}
	return result, nil
}
