package main

import (
	"fmt"
	"os"

	"usefull.pkg/des/internal/app/keys"
)

func main() {
	// Ask for the initial 8 ascii key
	var str string
	fmt.Println("What is you initial key: ")
	fmt.Scanln(&str)

	if len(str) != 8 {
		os.Exit(0)
	}

	keys := keys.GenerateKeys(str)
	for _, key := range keys {
		fmt.Println(key)
	}
}
