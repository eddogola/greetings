package greetings

import (
	"fmt"
	"math/rand"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string")
	}
	return fmt.Sprintf(greetingFormat(), name), nil
}

func Hellos(name string) (map[string]string, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}

	return map[string]string{
		name: fmt.Sprintf(greetingFormat(), name),
	}, nil
}

func greetingFormat() string {
	greetings := []string{
		"Hello, %v",
		"Nice to see you %v",
		"Most Welcome %v",
	}

	greeting := greetings[rand.Intn(len(greetings))]

	return greeting
}