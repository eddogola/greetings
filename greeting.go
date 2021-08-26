package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string")
	}
	return fmt.Sprintf("Hii, %v. Welcome!", name), nil
}
