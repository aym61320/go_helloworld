package main

import (
	"fmt"
	"main/slice"
)

func main() {

	msg := "Hello,World!"
	printer(msg)

	fmt.Println(returnFavorites())

	fmt.Println(returnCompliment("Hairstyle"))

	arrayPrinter()

	slice.SlicePrinter()
}

// Hello,World
func printer(msg string) {
	fmt.Println(msg)
}

// 複数return
func returnFavorites() (string, string) {
	return "EBI!", "Boston terrier!"
}

//変数return
func returnCompliment(word string) (msg string) {
	msg = "Good " + word
	return
}
