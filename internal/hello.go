package internal

import "fmt"

// SayHello ...
func SayHello(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// Size ...
func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}

	return "enormous"
}
