package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

func loadAPIKey() string {
	homeDir, _ := os.UserHomeDir()
	keyFilePath := homeDir + "/.gemini_key"

	// Check if the file "gemini_key" exists at ~/.gemini_key
	// If not, prompt user to input the key, create file and save the key in file
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		fmt.Print("Please enter your API key: ")
		bytePassword, _ := term.ReadPassword(0)
		key := string(bytePassword)
		if err := os.WriteFile(keyFilePath, []byte(key), 0600); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		return key
	}

	// Read the API key from file
	key, _ := os.ReadFile(keyFilePath)
	return strings.TrimSpace(string(key))
}

func main() {
	key := loadAPIKey()
	gemini := NewGemini(key, []map[string]interface{}{})
	for {
		fmt.Print("\n>>>>>>> ")
		var question string
		fmt.Scanln(&question)
		if question == "!BYE" {
			break
		}
		answer, ok := gemini.Ask(question)
		if !ok {
			fmt.Println("SYSTEM: Something went wrong")
		} else {
			fmt.Println("GEMINI:", answer)
		}
	}
}
