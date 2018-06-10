package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	store := make(map[string]bool)

	tmpStr := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter one string on one line.")
	for {
		tmpStr, _ = reader.ReadString('\n')
		tmpStr = strings.TrimSpace(tmpStr)
		if tmpStr == "" {
			break
		}

		if _, present := store[tmpStr]; present {
			fmt.Println("Key already present.")
		} else {
			store[tmpStr] = true
		}
	}

	fmt.Println("Key Store:")
	for key, _ := range store {
		fmt.Print(key, ".\n")
	}
	fmt.Println("\nBye.")
}
