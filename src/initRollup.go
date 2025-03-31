package src

import (
	"context"
	"fmt"
	"os"

	types "github.com/airchains-network/junction/x/rollup/types"
	"github.com/saatvik333/junction-go-client-template/utils"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func InitRollup() (string, error) {
	// Create Cosmos Client
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
	// var accountAddress = account.Address()
	chain_id := utils.ChainIdgenerate()
	da := utils.DaGen()
	keys, supply := utils.KeysGenerateAndSupply(3)

	// utils.Logger.Info("Initializing Rollup...")
	creator, _ := account.Address("air")
	fmt.Println(creator)

	// Create MsgInitRollup
	rollupMsg := &types.MsgInitRollup{
		Creator:                creator,
		Moniker:                "testing chain 11",
		ChainId:                chain_id,
		DenomName:              "airtoken",
		Keys:                   keys,
		Supply:                 supply,
		DaType:                 da,
		AclContractAddress:     "0x33c0B106c459d86841E96D58Db211Ae8554132d2",
		KmsVerifierAddress:     "0x61d1Ee49A472844985d7a8abd0FD482111de3389",
		TfheExecutorAddress:    "0x16054AEeDb074108193A3C074eb5e5B411577CD5",
		GatewayContractAddress: "0xc0665f2cD10beDC47ff04FEc169E55DB0B3BE77B",
		AscContractAddress:     "air1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqlnq2qd",
		RelayerGAddress:        "0x27CC9cA90057F3148AB6eAD6Dfd1Ff583fdFA67f",
		RelayerASCAddress:      "air1q8rztpy6cwcfl9pctahczdw6jneln7vdmmamdm",
	}

	// Broadcast transaction
	txResp, err := client.BroadcastTx(ctx, account, rollupMsg)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	fmt.Println("Rollup created successfully" + "Txhash " + txResp.TxHash)
	fmt.Println("http://localhost:26657/tx?hash=0x" + txResp.TxHash)

	rollupId := ""
	found := false
	for _, event := range txResp.Events {
		if event.Type == "rollup-initialized" {
			for _, attr := range event.Attributes {
				if attr.Key == "rollup-id" {
					rollupId = attr.Value
					found = true
				}
			}
		}
	}

	if !found {
		return "", fmt.Errorf("rollup ID not found but the attribute value exists")
	} else {
		if _, err := os.Stat("data"); os.IsNotExist(err) {
			err := os.Mkdir("data", 0755)
			if err != nil {
				return "", fmt.Errorf("failed to create data directory: %w", err)
			}
		}
		os.Create("data/rollupId.txt")
		os.WriteFile("data/rollupId.txt", []byte(rollupId), 0644)
	}

	fmt.Println("Rollup ID: " + rollupId)
	return txResp.TxHash, nil
}
