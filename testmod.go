package testmod

// https://roberto.selbach.ca/intro-to-go-modules/

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
