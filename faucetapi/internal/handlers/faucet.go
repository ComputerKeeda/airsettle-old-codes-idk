package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	utils "github.com/airchains-network/settlementfaucet/internal/utils"
	// "github.com/airchains-network/settlementfaucet/internal/models"
)

type accountDetails struct {
	Address string `json:"address"`
}

type jsonResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transactionhash"`
	Message         string `json:"message"`
	Description     string `json:"description"`
}

func FaucetHandler(w http.ResponseWriter, r *http.Request) {

	// Create the response object
	finalResponse := jsonResponse{
		Success:         false,
		TransactionHash: "",
		Message:         "",
		Description:     "",
	}

	var input accountDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	recipientAddress := input.Address

	// check if address is valid
	isvalid := utils.IsValidCosmosAddress(recipientAddress)
	if !isvalid {
		http.Error(w, "Invalid address", http.StatusBadRequest)
		return
	}

	// check balance, send if balance < 10 tokens
	const (
		accountName   = "alice"                         // "vitalikbhoot"
		path          = "~/.Airchains-settlement-layer" // "./accounts"
		addressPrefix = "air"
		gasLimit      = "3000000"
		blockchainRPC = "http://0.0.0.0:1317"
		tendermintRPC = "http://localhost:26657"
	)
	ctx := context.Background()
	client, err := cosmosclient.New(ctx, cosmosclient.WithGas(gasLimit), cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(tendermintRPC), cosmosclient.WithKeyringDir(path))
	if err != nil {
		http.Error(w, "Blockchain connection error", http.StatusBadRequest)
	}

	// get balance
	balanceAmount, err := client.BankBalances(ctx, recipientAddress, nil)
	maxTokens := uint64(10)

	for _, balance := range balanceAmount {
		if balance.Denom == "token" {
			tokenAmount := balance.Amount
			tokenint := tokenAmount.Uint64()
			if tokenint > maxTokens {
				finalResponse = jsonResponse{
					Success:         false,
					TransactionHash: "",
					Message:         "Balance > 10 tokens",
					Description:     "Already have enough balance: " + string(tokenint) + balance.Denom + ", so can not send more",
				}
				responseJSON, err := json.Marshal(finalResponse)
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(responseJSON)
				return
			}
		}
	}

	// send tokens
	amountToSend := uint(10)
	account, err := client.Account(accountName)
	if err != nil {
		finalResponse = jsonResponse{
			Success:         false,
			TransactionHash: "",
			Message:         "Faucet type error",
			Description:     string(err.Error()),
		}
		responseJSON, err := json.Marshal(finalResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}

	amountWithDenom := fmt.Sprintf("%dtoken", amountToSend)
	coins, err := sdk.ParseCoinsNormalized(amountWithDenom)
	if err != nil {
		finalResponse = jsonResponse{
			Success:         false,
			TransactionHash: "",
			Message:         "Failed to parse coins",
			Description:     string(err.Error()),
		}
		responseJSON, err := json.Marshal(finalResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}

	sendTx, err := client.BankSendTx(ctx, account, recipientAddress, coins)
	if err != nil {
		finalResponse = jsonResponse{
			Success:         false,
			TransactionHash: "",
			Message:         "Transaction Formation error",
			Description:     string(err.Error()),
		}
		responseJSON, err := json.Marshal(finalResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}

	txResponse, err := sendTx.Broadcast(ctx)
	if err != nil {
		finalResponse = jsonResponse{
			Success:         false,
			TransactionHash: "",
			Message:         "Transaction Failed",
			Description:     "Faucet empty OR faucet/blockchain conjusted" + string(err.Error()),
		}
		responseJSON, err := json.Marshal(finalResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	} else {
		finalResponse = jsonResponse{
			Success:         true,
			TransactionHash: txResponse.TxHash,
			Message:         "Transaction successful",
			Description:     "Can check txhash on:" + blockchainRPC + "/cosmos/tx/v1beta1/txs/" + txResponse.TxHash + " and balance on: " + blockchainRPC + "/cosmos/bank/v1beta1/spendable_balances/" + recipientAddress,
		}

		responseJSON, err := json.Marshal(finalResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}
}
