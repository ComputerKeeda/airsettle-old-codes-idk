package verifier

import (
	"encoding/json"
	"fmt"
	"math/big"

	groth16 "github.com/arnaucube/go-snark/groth16"
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

func Verifier(zkproof string, vkey string, input string) bool {
	// zkproof := `{"pi_a": ["123", "456", "789"], ...}`
	proof, err := GetProof(zkproof)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// vkey := `{"protocol": "groth16", ...}`
	vk, err := GetVerificationKey(vkey)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// input := `["123", "456", "789"]`
	inputs, err := GetInputs(input)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	res := groth16.VerifyProof(vk, proof, inputs, true)
	return res
}

func GetProof(proofData string) (groth16.Proof, error) {
	var zkproof Proof
	err := json.Unmarshal([]byte(proofData), &zkproof)
	if err != nil {
		return groth16.Proof{}, err
	}

	// Convert string values to big.Int
	var piA [3]*big.Int
	for i, value := range zkproof.PiA {
		intVal, success := new(big.Int).SetString(value, 10)
		if !success {
			return groth16.Proof{}, fmt.Errorf("failed to parse pi_a value: %s", value)
		}
		piA[i] = intVal
	}

	var piB [3][2]*big.Int
	for i, values := range zkproof.PiB {
		if len(values) != 2 {
			return groth16.Proof{}, fmt.Errorf("pi_b should have 2 values in each row")
		}
		for j, value := range values {
			intVal, success := new(big.Int).SetString(value, 10)
			if !success {
				return groth16.Proof{}, fmt.Errorf("failed to parse pi_b value: %s", value)
			}
			piB[i][j] = intVal
		}
	}

	var piC [3]*big.Int
	for i, value := range zkproof.PiC {
		intVal, success := new(big.Int).SetString(value, 10)
		if !success {
			return groth16.Proof{}, fmt.Errorf("failed to parse pi_c value: %s", value)
		}
		piC[i] = intVal
	}

	// Create and return the proof
	sampleProof := groth16.Proof{
		PiA: piA,
		PiB: piB,
		PiC: piC,
	}

	return sampleProof, nil
}

func GetVerificationKey(verificationKeyData string) (groth16.Vk, error) {
	var vkey VerificationKey
	err := json.Unmarshal([]byte(verificationKeyData), &vkey)
	if err != nil {
		return groth16.Vk{}, err
	}

	var samplevk groth16.Vk

	ic := make([][3]*big.Int, len(vkey.IC))
	for i := 0; i < len(vkey.IC); i++ {
		ic[i][0], _ = new(big.Int).SetString(vkey.IC[i][0], 10)
		ic[i][1], _ = new(big.Int).SetString(vkey.IC[i][1], 10)
		ic[i][2], _ = new(big.Int).SetString(vkey.IC[i][2], 10)
	}
	samplevk.IC = ic

	vkAlpha1 := [3]*big.Int{}
	for i := 0; i < 3; i++ {
		vkAlpha1[i], _ = new(big.Int).SetString(vkey.VKAlpha1[i], 10)
	}
	samplevk.G1.Alpha = vkAlpha1

	vkBeta2 := [3][2]*big.Int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			vkBeta2[i][j], _ = new(big.Int).SetString(vkey.VKBeta2[i][j], 10)
		}
	}
	samplevk.G2.Beta = vkBeta2

	vkGamma2 := [3][2]*big.Int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			vkGamma2[i][j], _ = new(big.Int).SetString(vkey.VKGamma2[i][j], 10)
		}
	}
	samplevk.G2.Gamma = vkGamma2

	vkDelta2 := [3][2]*big.Int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			vkDelta2[i][j], _ = new(big.Int).SetString(vkey.VKDelta2[i][j], 10)
		}
	}
	samplevk.G2.Delta = vkDelta2

	return samplevk, nil
}

func GetInputs(inputs string) ([]*big.Int, error) {
	var input []*json.RawMessage
	err := json.Unmarshal([]byte(inputs), &input)
	if err != nil {
		return nil, err
	}

	// Convert JSON raw messages to big.Int
	sampleInputs := make([]*big.Int, len(input))
	for i, rawMessage := range input {
		var valueStr string
		err := json.Unmarshal(*rawMessage, &valueStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input value: %w", err)
		}

		intVal, success := new(big.Int).SetString(valueStr, 10)
		if !success {
			return nil, fmt.Errorf("failed to parse input value: %s", valueStr)
		}
		sampleInputs[i] = intVal
	}

	return sampleInputs, nil
}
