package handlers

import (
	utils "github.com/ComputerKeeda/airsettle/client_helper/utils"

	cosmosAccount "github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

/*
CreateAccount creates a new account using the provided path and account name.
It uses the Cosmos Account registry to create the account.

Parameters:

- path: The home path for the Cosmos Account registry.
- accountName: The name of the account to be created.

Returns:
- response: A utils.Response containing the status and message of the account creation process.

Example:

	package main

	import (
		handlers "github.com/ComputerKeeda/airsettle_Calls/handlers"
	)

	func main() {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		var newDir = wd + "/test"
		res := handlers.CreateAccount(newDir,"test")
		fmt.Println(res.Message)
	}
*/
func CreateAccount(path string, accountName string) (response utils.Response) {

	registry, err := cosmosAccount.New(cosmosAccount.WithHome(path))

	if err != nil {
		panic(err)
	}

	_, mnemonic, err2 := registry.Create(accountName)
	if err2 != nil {
		panic(err)
	}

	return utils.Response{
		Status:  true,
		Message: "Account created with mnemonic: " + mnemonic,
	}

}

// func unicodeToDecimalString(input string) string {
// 	decimalStrings := []string{}
// 	for _, char := range input {
// 		decimalStrings = append(decimalStrings, strconv.Itoa(int(char)))
// 	}
// 	return "[" + strings.Join(decimalStrings, ", ") + "]"
// }
