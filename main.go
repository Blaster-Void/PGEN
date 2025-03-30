// this package for generating random characters and save to file in Document
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
) // need to import

func saveToJson(content password) error { // Save password to json file
	homeDir, err := os.UserHomeDir() // Get home directory
	if err != nil {
		return err
	}
	file, err := os.Create(homeDir + "/mypassword.json") // Create json file
	if err != nil {
		return err
	}
	defer file.Close() // Close file

	jsonContent := fmt.Sprintf(`pass: %s`, content.Pass) // Create json content
	_, err = file.WriteString(jsonContent)               // Write json content to file
	if err != nil {
		return err
	}

	return nil
}

func ShowPasswords() error { // Show password from json file
	homeDir, err := os.UserHomeDir() // Get home directory
	if err != nil {
		return err
	}
	file, err := os.Open(homeDir + "/mypassword.json") // Open json file
	if err != nil {
		return err
	}
	defer file.Close() // Close file

	content, err := ioutil.ReadAll(file) // Read json content
	if err != nil {
		return err
	}

	fmt.Println(string(content)) // Print json content

	return nil
}

type password struct { // Define password struct
	Pass string `json:"pass"`
}

// 65 ta 122 a Z
func main() { // Define main function

	var x int                                   // Initialize x for random number generation
	if len(os.Args) == 2 || len(os.Args) == 3 { // Check if arguments are provided
		if os.Args[1] == "-c" { // Check if command is -c
			var allchars []rune          // Initialize allchars slice
			op := os.Args[2]             // Get operation argument
			opint, _ := strconv.Atoi(op) // Convert operation argument to integer
			for i := 0; i < opint; i++ {

				x = rand.Intn(58) + 65               // Generate random number between 65 and 122
				allchars = append(allchars, rune(x)) // Append random character to allchars slice
				fmt.Print(string(rune(x)))           // Print random character

				time.Sleep(time.Millisecond * 60) // Sleep for 60 milliseconds for decoration

			}
			saveToJson(password{Pass: string(allchars)}) // Save password to json file
			fmt.Println("")                              // Print newline after saving password
		} else if os.Args[1] == "-show" { // Check if command is -show
			ShowPasswords() // Show passwords
		}
	} else { // Check if no arguments are provided
		var allchars []rune // Initialize allchars slice
		for i := 0; i < 8; i++ {

			x = rand.Intn(58) + 65               // Generate random number between 65 and 122
			allchars = append(allchars, rune(x)) // Append random character to allchars slice
			fmt.Print(string(rune(x)))           // Print random character

			time.Sleep(time.Millisecond * 60) // Sleep for 60 milliseconds for decoration

		}
		fmt.Println("\ndone")                        // Print newline after done message
		fmt.Println("")                              // Print newline after done message
		saveToJson(password{Pass: string(allchars)}) // Save password to json file

	}
}
