package testmod

import (
	"fmt"
)

func SayHello(name, str string) string {
	return fmt.Sprintf("Hi, welcome %s, %s", name, str)
}
