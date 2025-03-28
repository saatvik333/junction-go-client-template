package junction

import (
	"context"
	"fmt"

	"github.com/saatvik333/junction-go-client-template/utils"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func ClientConnect(accountPath, accountName, addressPrefix, jsonRpc string) (Account cosmosaccount.Account, Addr string, client cosmosclient.Client) {

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Error creating account registry: %v", err))
		return Account, Addr, client
	}

	Account, err = registry.GetByName(accountName)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Error getting account: %v", err))
		return Account, Addr, client
	}

	Addr, err = Account.Address(addressPrefix)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Error getting address: %v", err))
		return Account, Addr, client
	}

	ctx := context.Background()
	client, err = cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(jsonRpc), cosmosclient.WithHome(accountPath), cosmosclient.WithGas("auto"))
	if err != nil {
		utils.Log.Error("Switchyard client connection error")
		utils.Log.Error(err.Error())
		return Account, Addr, client
	}
	return Account, Addr, client
}
