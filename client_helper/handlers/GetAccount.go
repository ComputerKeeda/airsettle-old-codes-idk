package handlers

import (
	"fmt"
	"log"

	utils "airsettle/client_helper/utils"
	cosmosAccount "github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

func GetAccount(path string, accountName string) (response utils.Response) {
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

	data, err := account.Record.Marshal()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", data)

	return utils.Response{
		Status:  true,
		Message: fmt.Sprintf("Account address: %s", addr),
	}
}
