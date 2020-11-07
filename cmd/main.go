package main

import (
	"fmt"
	"os"

	"go-des/internal/app/encryption"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 1 && os.Args[1] == "--help" {
		// print the --help page
		helper()
	} else if len(os.Args) > 1 && os.Args[1] == "-e" && os.Args[2] != "" {
		// if -b64 -> encode in base64
		// else -h hexa
		// else binary
		if os.Args[2] == "-b64" {
			encryption.EncryptMessage(os.Args[3], 1)
		} else if os.Args[2] == "-h" {
			encryption.EncryptMessage(os.Args[3], 2)
		} else if os.Args[2] == "-b" {
			encryption.EncryptMessage(os.Args[3], 0)
		} else {
			encryption.EncryptMessage(os.Args[2], 0)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "-d" && os.Args[2] != "" {
		// if -b64 -> decode in base64
		// else -h hexa
		// else binary
		if os.Args[2] == "-b64" {
			encryption.DecryptMessage(os.Args[3], 1)
		} else if os.Args[2] == "-h" {
			encryption.DecryptMessage(os.Args[3], 2)
		} else if os.Args[2] == "-b" {
			encryption.DecryptMessage(os.Args[3], 0)
		} else {
			encryption.DecryptMessage(os.Args[2], 0)
		}
	} else {
		// print the helper is no arguments is given
		helper()
	}
}

func helper() {
	fmt.Println("===================================================")
	fmt.Println("Welcome to go-DES")
	fmt.Println("===================================================")
	fmt.Println("Example:")
	fmt.Print("\n")
	fmt.Println("$ go-DES -e \"I am a message to encrypt\": Will encrypt the message in binary")
	fmt.Println("$ go-DES -e -b \"I am a message to encrypt\": Will encrypt the message in binary")
	fmt.Println("$ go-DES -e -h \"I am a message to encrypt\": Will encrypt the message in hexadecimal")
	fmt.Println("$ go-DES -e -b64 \"I am a message to encrypt\": Will encrypt the message in base64")
	fmt.Println("$ go-DES -d -h \"762336c2b5360cc38751\": Will decrypt the hexadecimal message")
	fmt.Println("===================================================")
	fmt.Println("Arguments:")
	fmt.Print("\n")
	fmt.Println("--help: Print this helper")
	fmt.Println("-e: Encrypt a message")
	fmt.Println("-d: Decrypt a message")
	fmt.Println("-h: Use hexadecimal")
	fmt.Println("-b64: Use base64")
	fmt.Println("-b: Use binary")
	fmt.Println("===================================================")
}
