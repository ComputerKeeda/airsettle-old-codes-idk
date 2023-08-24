package utils 

import (
	"strings"
	"github.com/cosmos/btcutil/bech32"
)

func IsValidCosmosAddress(address string) bool {
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