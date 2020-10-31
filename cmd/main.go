package main

import (
	"fmt"
	"os"
	"strings"

	"usefull.pkg/des/internal/app/keys"
	"usefull.pkg/des/internal/app/rounds"
	"usefull.pkg/des/internal/pkg/binary"
	"usefull.pkg/des/internal/pkg/permutations"
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

func main() {
	var word string
	fmt.Println("Insert a 8 character word: ")
	fmt.Scanln(&word)

	// convert the message into binary
	b := binary.StringToBinary(word)
	// split the binary string into a slice of strings
	m := strings.Split(b, "")
	fmt.Println("binary: ")
	fmt.Println(m)

	// make the IP inital permutation
	fmt.Println("permutation:")
	ti := permutations.Ip(m)
	fmt.Println(ti)

	// generate the keys for the rounds
	keys := getKeys()

	// make the rounds
	for i, k := range keys {
		fmt.Println(rounds.Round(ti, k, i))
		break
	}

}
