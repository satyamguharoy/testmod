package testmod

import "fmt"

// Hi handles greetings
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
