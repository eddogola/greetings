package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hii, %v. Welcome!", name)
}
