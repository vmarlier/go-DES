package main

import (
	"fmt"
	"os"

	"go-des/internal/app/encryption"
)

func encrypt(message string) {
	encryption.EncryptMessage(message)
}

func decrypt(message string) {
	encryption.DecryptMessage(message)
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 1 && os.Args[1] == "-h" {
		helper()
	} else if len(os.Args) > 1 && os.Args[1] == "-e" && os.Args[2] != "" {
		encrypt(os.Args[2])
	} else if len(os.Args) > 1 && os.Args[1] == "-d" && os.Args[2] != "" {
		decrypt(os.Args[2])
	} else {
		helper()
	}
}

func helper() {
	fmt.Println("===================================================")
	fmt.Println("Welcome to go-DES")
	fmt.Println("===================================================")
	fmt.Println("Example:")
	fmt.Println("go-DES -e \"I am a message to encrypt\"")
	fmt.Println("===================================================")
	fmt.Println("Arguments:")
	fmt.Print("\n")
	fmt.Println("-h: Print this helper")
	fmt.Println("-e: Encrypt a message")
	fmt.Println("-d: Decrypt a message")
	fmt.Println("===================================================")
}
