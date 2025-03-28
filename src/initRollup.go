package junction

import (
	types "github.com/airchains-network/junction/x/rollup/types"
	RandomGen "github.com/saatvik333/junction-go-client-template/utils"
	// "github.com/saatvik333/junction-go-client-template/utils"
)

func InitRollup(creator string) string {

	moniker := RandomGen.Monikergenerate()
	da:= RandomGen.DaGen()
	keys, supply := RandomGen.KeysGenerateAndSupply(3)

	msg := &types.MsgInitRollup{
		Creator: creator,
		Moniker: moniker,
		ChainId: "test-chain",
		DenomName: "test",
		DaType: da,
		Keys: keys,
		Supply : supply,
	}
	
	return 
}

