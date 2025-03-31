package junction

import (
	"context"
	"fmt"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func ClientConnect(accountPath, accountName, addressPrefix, jsonRpc string) (Account cosmosaccount.Account, Addr string, client cosmosclient.Client) {

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		// fmt.Printf("Error creating account registry: %v", err))
		fmt.Printf("Error creating account registry: %v", err)
		return Account, Addr, client
	}

	Account, err = registry.GetByName(accountName)
	if err != nil {
		fmt.Printf("Error getting account: %v", err)
		return Account, Addr, client
	}

	Addr, err = Account.Address(addressPrefix)
	if err != nil {
		fmt.Printf("Error getting address: %v", err)
		return Account, Addr, client
	}

	ctx := context.Background()
	client, err = cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(jsonRpc), cosmosclient.WithHome(accountPath), cosmosclient.WithGas("auto"))
	if err != nil {
		fmt.Printf("SwitchyardPlient connection  error %v",err)
		return Account, Addr, client
	}
	return Account, Addr, client
}
