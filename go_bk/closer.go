package main

import "fmt"

// クロージャ
func Later() func(string) string {
	// storeは破棄されない
    var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("World"))
	fmt.Println(f("From"))
	fmt.Println(f("Golang"))
}
