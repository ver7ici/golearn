package main

import (
	"fmt"

	"github.com/ver7ici/golearn/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
