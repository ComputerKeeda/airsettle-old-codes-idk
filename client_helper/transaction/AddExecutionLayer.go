package transaction

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Airchains-Studio/Settlement_Layer/x/airsettle/types"

	"github.com/cosmos/btcutil/bech32"
	cosmosAccount "github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosClient "github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddExecutionLayer(name string, path string) {

	// MsgAddExecutionLayer require  more gas fee
	gasLimit := "2000000"

	ctx := context.Background()
	addressPrefix := "air"

	// Create a Cosmos client instance
	// tmpDir := "/tmp/airchain"
	client, err := cosmosClient.New(ctx, cosmosClient.WithGas(gasLimit), cosmosClient.WithAddressPrefix(addressPrefix), cosmosClient.WithNodeAddress("http://localhost:26657"))
	if err != nil {
		log.Fatal(err)
	}

	// if !isValidCosmosAddress(accountAddress) {
	// 	fmt.Println("Invalid NewValidatorAddress")
	// 	return
	// }

	registry, err := cosmosAccount.New(cosmosAccount.WithHome(path))
	if err != nil {
		panic(err)
	}

	account, err := registry.GetByName(name)
	if err != nil {
		panic(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	// accountName := "Isabella"

	// // Get account from the keyring
	// account, err := client.Account(accountAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account)
	// os.Exit(0)

	// addr, err := account.Address(addressPrefix)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// data, err := os.ReadFile("json/verification_key.json")
	data, err := os.ReadFile("./sample_json_data/verification_key.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Convert the data to a string
	verificationKeyString := string(data)

	msg := &types.MsgAddExecutionLayer{
		Creator:         addr,
		VerificationKey: verificationKeyString,
		ChainInfo:       "This chain is build by airchains for testing purposes. copyright 2023",
	}
	_ = msg

	// http://0.0.0.0:1317/cosmos/tx/v1beta1/txs/{transaction_hash}
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(txResp)
	fmt.Println("http://0.0.0.0:1317/cosmos/tx/v1beta1/txs/" + txResp.TxHash)
}

func isValidCosmosAddress(address string) bool {
	const customPrefix = "air"
	// Check if the address has the correct prefix
	if !strings.HasPrefix(address, customPrefix) {
		return false
	}
	// Decode the Bech32 encoded address
	_, _, err := bech32.Decode(address, bech32.MaxLengthBIP173)

	// Return true if decoding was successful, false otherwise
	return err == nil
}
