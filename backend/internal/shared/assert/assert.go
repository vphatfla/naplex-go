package assert

import "fmt"

// Assert the type T, return the value with type T and error
func AssertAndGetValue[T any](value any) (T, error) {
	v, ok := value.(T)
	if !ok {
		var temp T
		return temp, fmt.Errorf("value %v type %T but required type %T", v, v, temp)
	}
	return v, nil
}
