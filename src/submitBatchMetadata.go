package src

import (
	"context"
	"fmt"

	types "github.com/airchains-network/junction/x/rollup/types"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"github.com/saatvik333/junction-go-client-template/utils"

	
)

func MsgSubmitBatchMetadata ()(string,error){
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
	da:= utils.DaGen()

	BatchMetadatamsg := &types.MsgSubmitBatchMetadata{
		Creator:      creator,
		BatchNo:      2,
		RollupId:     "55f8fb2fffc2632450dc2797ca4e33f39d239794d40bddf32b0ed08ba2d8ce06",
		DaName:       da,
		DaCommitment: "wwgdahdhadh23e743ewqd34rwdewdw",
		DaHash:       "ggyggy554537gfddhhuu35etdgchfhnvgf",
		DaPointer:    "fft6t33u9u577tffgfgygygye3465ttr",
		DaNamespace:  "tftfvgwugf4667rccsfovrddressgs66vvzs",
	}
	txBatchResp , err := client.BroadcastTx(ctx,account,BatchMetadatamsg)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println("successfull batchMetadata"+ "TxHash ", txBatchResp.TxHash)

	return txBatchResp.TxHash , nil 
}