package testmod

import (
	"fmt"
)

func SayHello(name, str string) string {
	return fmt.Sprintf("Hi, this is github.com, welcome %s, %s", name, str)
}
