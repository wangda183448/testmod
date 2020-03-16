package testmod

import (
	"fmt"
)

func SayHello(name, str string) string {
	return fmt.Sprintf("Hi, github welcome %s, %s", name, str)
}
