package main

import (
    "fmt"

    "github.com/kekcleader/go-modules/greetings"
)

func main() {
    message := greetings.Hello("Gladiolus")
    fmt.Println(message)
}
