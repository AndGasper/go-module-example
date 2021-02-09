package main

import "fmt"
import "moduleone"

func main() {
	message := moduleone.ModuleOne("Some Name")
	fmt.Println(message)
}