package main

import (
	"fmt"

	"go-des/internal/app/encryption"
)

func encrypt() {
	var word string
	fmt.Println("Insert a 8 character word: ")
	fmt.Scanln(&word)

	encryption.EncryptMessage(word)
}

func main() {
	encrypt()
}
