package src

import (
	"context"
	"fmt"

	types "github.com/airchains-network/junction/x/rollup/types"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func MsgSubmitBatch ()(string,error){
	ctx := context.Background()
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("air"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	account, err := client.Account("alice") 
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	creator,_ := account.Address("air")

	Batchmsg := &types.MsgSubmitBatch{
		Creator:                creator,
		RollupId:               "55f8fb2fffc2632450dc2797ca4e33f39d239794d40bddf32b0ed08ba2d8ce06",
		BatchNo:               1,
		MerkleRootHash:         "abcdeeeeee1",
		PreviousMerkleRootHash: "abcheeeeee0",
		ZkProof:                []byte{0xee, 0xf5, 0xd8, 0x76, 0x23}, 
		PublicWitness:          []byte{0x77, 0x65, 0x68, 0x77, 0x64},
	}
	txBatchResp , err := client.BroadcastTx(ctx,account,Batchmsg)
	if err != nil {
		fmt.Errorf(err.Error())
		return "", err
	}
	fmt.Println("successfull batch submission"+ "TxHash ", txBatchResp.TxHash)

	return txBatchResp.TxHash , nil 
}