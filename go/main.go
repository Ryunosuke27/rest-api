package main

import (
    "fmt"
    "go/config"
)

func main() {
    fmt.Println(config.Config.Port)
    fmt.Println("hello")
}
