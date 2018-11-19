// https://roberto.selbach.ca/intro-to-go-modules/
package testmod

import (
	"errors"
	"fmt"
)

// Hi returns a friendly greeting
func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "cn":
		return fmt.Sprintf("你好, %s", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
