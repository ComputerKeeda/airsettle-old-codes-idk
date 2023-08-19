package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	handlers "airsettle/client_helper/handlers"

	randomData "github.com/Pallinder/go-randomdata"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var newDir = wd + "/storage_account"
	fmt.Printf("%v", newDir)
	randomName := randomData.FirstName(randomData.RandomGender)
	data := handlers.CreateAccount(newDir, randomName)
	PrettyPrint(data)

	data2 := handlers.GetAccount(newDir, randomName)
	PrettyPrint(data2)

}

// The PrettyPrint function takes in any data and prints it out in a formatted and indented JSON
// string.
func PrettyPrint(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		fmt.Printf("Failed to generate JSON: %v", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

// The function `InputFileToTextConverter` takes a file path as input, reads the contents of the file,
// and returns the contents as a slice of strings.

func InputFileToTextConverter(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []string
	err = json.NewDecoder(file).Decode(&inputs)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}
