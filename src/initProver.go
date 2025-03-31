package src

import (
	"context"
	"fmt"
	"log"
	"os"

	types "github.com/airchains-network/junction/x/rollup/types"
	"github.com/saatvik333/junction-go-client-template/utils"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func loadVerificationKey(filename string) ([]byte, error) {
	// Read the JSON file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return fileData, nil
}
func InitProver() (string, error) {
	ctx := context.Background()
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("air"), cosmosclient.WithGas("auto"), cosmosclient.WithGasAdjustment(1.5))
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

	creator, _ := account.Address("air")

	rollupId, err := os.ReadFile("data/rollupId.txt")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	proverMsg := &types.MsgInitProver{
		Creator:               creator,
		RollupId:              string(rollupId),
		ProverVerificationKey: proverKey,
		ProverType:            "abfg",
		ProverEndpoint:        "test",
	}
	txResp, err := client.BroadcastTx(ctx, account, proverMsg)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	utils.Logger.Info("Rollup Prover Initialized, Transaction Hash:" + txResp.TxHash)
	utils.Logger.Info("http://localhost:26657/tx?hash=0x" + txResp.TxHash)
	utils.Logger.Success("http://localhost:1317/airchains-network/junction/rollup/get_rollup_info/" + string(rollupId))
	return txResp.TxHash, nil
}
