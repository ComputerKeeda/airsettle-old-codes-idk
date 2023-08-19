package handlers

import (
	"fmt"
	"log"

	utils "github.com/Airchains-Studio/Settlement_Layer/client_helper/utils"
	cosmosAccount "github.com/ignite/cli/ignite/pkg/cosmosaccount"
)


/*
GetAccount retrieves an existing account using the provided path and account name from the Cosmos Account registry.

Parameters:

- path: The home path for the Cosmos Account registry.
- accountName: The name of the account to be retrieved.

Returns:
- response: A utils.Response containing the status and details of the requested account.

Example:

	package main

	import (
		handlers "github.com/Airchains-Studio/Settlement_Layer_Calls/handlers"
	)

	func main() {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		var searchDir = wd + "/test"
		res := handlers.GetAccount(searchDir, "test")
		if res.Status {
			fmt.Println("Account found:", res.Message)
		} else {
			fmt.Println("Account not found:", res.Message)
		}
	}
*/
func GetAccount(path string, accountName string) (response utils.AccountResponse) {
	addressPrefix := "air"
	registry, err := cosmosAccount.New(cosmosAccount.WithHome(path))
	if err != nil {
		panic(err)
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		panic(err)
	}


	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// data, err := account.Record.Marshal()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return utils.AccountResponse{
		Status:  true,
		Message: fmt.Sprintf("Account address: %s", addr),
		Address: addr,
	}
}
