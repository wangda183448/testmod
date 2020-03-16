package testmod

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hi, welcome %s", name)
}
