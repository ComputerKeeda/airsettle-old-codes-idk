package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	groth16 "github.com/arnaucube/go-snark/groth16"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Proof struct {
	PiA      []string   `json:"pi_a"`
	PiB      [][]string `json:"pi_b"`
	PiC      []string   `json:"pi_c"`
	Protocol string     `json:"protocol"`
	Curve    string     `json:"curve"`
}

type VerificationKey struct {
	Protocol      string          `json:"protocol"`
	Curve         string          `json:"curve"`
	NPublic       int             `json:"nPublic"`
	VKAlpha1      [3]string       `json:"vk_alpha_1"`
	VKBeta2       [3][2]string    `json:"vk_beta_2"`
	VKGamma2      [3][2]string    `json:"vk_gamma_2"`
	VKDelta2      [3][2]string    `json:"vk_delta_2"`
	VKAlphaBeta12 [2][3][2]string `json:"vk_alphabeta_12"`
	IC            [][3]string     `json:"IC"`
}

// ! Utility Functions are here --> Start

func convertToVk(vk VerificationKey) groth16.Vk {
	var result groth16.Vk

	// Convert IC
	for _, ic := range vk.IC {
		var icConverted [3]*big.Int
		for j, value := range ic {
			icConverted[j] = big.NewInt(0)
			icConverted[j].SetString(value, 10)
		}
		result.IC = append(result.IC, icConverted)
	}

	// Convert VKAlpha1
	for i, value := range vk.VKAlpha1 {
		result.G1.Alpha[i] = big.NewInt(0)
		result.G1.Alpha[i].SetString(value, 10)
	}

	// Convert VKBeta2, VKGamma2, VKDelta2
	convertToArray := func(arr [3][2]string) [3][2]*big.Int {
		var result [3][2]*big.Int
		for i, subArr := range arr {
			for j, value := range subArr {
				result[i][j] = big.NewInt(0)
				result[i][j].SetString(value, 10)
			}
		}
		return result
	}

	result.G2.Beta = convertToArray(vk.VKBeta2)
	result.G2.Gamma = convertToArray(vk.VKGamma2)
	result.G2.Delta = convertToArray(vk.VKDelta2)

	return result
}

func convertToProof(original Proof) groth16.Proof {
	var result groth16.Proof

	// Convert PiA
	for i, value := range original.PiA {
		result.PiA[i] = big.NewInt(0)
		result.PiA[i].SetString(value, 10)
	}

	// Convert PiB
	for i, subArr := range original.PiB {
		for j, value := range subArr {
			result.PiB[i][j] = big.NewInt(0)
			result.PiB[i][j].SetString(value, 10)
		}
	}

	// Convert PiC
	for i, value := range original.PiC {
		result.PiC[i] = big.NewInt(0)
		result.PiC[i].SetString(value, 10)
	}

	return result
}

func stringToBigInts(inputJSON string) ([]*big.Int, error) {
	var stringNumbers []string
	err := json.Unmarshal([]byte(inputJSON), &stringNumbers)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %w", err)
	}

	bigInts := make([]*big.Int, len(stringNumbers))
	for i, stringNumber := range stringNumbers {
		bigInts[i] = big.NewInt(0)
		_, success := bigInts[i].SetString(stringNumber, 10)
		if !success {
			return nil, fmt.Errorf("Failed to convert string %q to big.Int", stringNumber)
		}
	}

	return bigInts, nil
}

// ! Utility Functions are here --> End

func (k Keeper) Verifier(ctx sdk.Context, zkproof string, vkey string, input string) (bool, string) {

	_ = ctx

	// formate verification key
	var semiParsedVerficationKey VerificationKey
	err := json.Unmarshal([]byte(vkey), &semiParsedVerficationKey)
	if err != nil {
		return false, "ERROR: semiParsedVerficationKey"
	}
	parsedVerficationKey := convertToVk(semiParsedVerficationKey)

	// formate proof
	var semiParsedProof Proof
	err2 := json.Unmarshal([]byte(zkproof), &semiParsedProof)
	if err2 != nil {
		return false, "Error: semiParsedProof"
	}
	parsedProof := convertToProof(semiParsedProof)

	// formate input
	parsedInputs, err := stringToBigInts(input)
	if err != nil {
		return false, "Error: parsedInputs"
	}

	res := groth16.VerifyProof(parsedVerficationKey, parsedProof, parsedInputs, true)
	return res, "verification successful"
}
