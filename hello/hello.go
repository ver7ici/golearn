package main

import (
	"fmt"

	"github.com/ver7ici/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
