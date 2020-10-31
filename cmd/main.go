package main

import (
	"fmt"
	"os"

	"usefull.pkg/des/internal/app/keys"
)

func getKeys() []string {
	// Ask for the initial 8 ascii key
	var str string
	fmt.Println("What is you initial key: ")
	fmt.Scanln(&str)

	if len(str) != 8 {
		os.Exit(0)
	}

	return keys.GenerateKeys(str)
}

func encrypt() {
	var word string
	fmt.Println("Insert a 8 character word: ")
	fmt.Scanln(&word)

	//return encrypt.EncryptMessage(word)
}

func main() {
	encrypt()
}
