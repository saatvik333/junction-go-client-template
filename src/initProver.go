package src

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	types "github.com/airchains-network/junction/x/rollup/types"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

type VerificationData struct {
	ProverVerificationKey []byte `json:"proverVerificationKey,omitempty"`
}

func loadVerificationKey(filename string) ([]byte, error) {
	// Read the JSON file
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse JSON into struct
	var jsonData VerificationData
	err = json.Unmarshal(fileData, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return jsonData.ProverVerificationKey, nil
}
func InitProver() (string,error){
	ctx := context.Background()
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("air"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Get account from keyring
	account, err := client.Account("alice") 
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	proverKey, err := loadVerificationKey("verification.json")
	if err != nil {
		log.Fatal("Failed to load verification key:", err)
	}

	// var accountAddress = account.Address()

	// utils.Logger.Info("Initializing Rollup...")
	creator,_ := account.Address("air")

	proverMsg := &types.MsgInitProver{
		Creator: creator,
		RollupId:"55f8fb2fffc2632450dc2797ca4e33f39d239794d40bddf32b0ed08ba2d8ce06",
		ProverVerificationKey: proverKey,
		ProverType:  "abfg"         ,
		ProverEndpoint:  "test"      ,

	}
	txResp, err := client.BroadcastTx(ctx, account, proverMsg)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	fmt.Println("Proved " + txResp.TxHash)
	return  txResp.TxHash ,nil 
}