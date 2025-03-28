package utils

import (
	"fmt"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"encoding/json"	
	"os"
)

// AccountData stores account connection data.
type AccountData struct {
	Account cosmosaccount.Account
	Addr    string
	Client  cosmosclient.Client
}


func CheckIfAccountExists(accountName string, client cosmosclient.Client, addressPrefix string, accountPath string) (bool, string) {

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		return false, ""
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		fmt.Println("Failed to get the Address:", err)
		return false, ""
	}

	return true, addr
}

func FetchAccount(accountName string, client cosmosclient.Client, addressPrefix string, accountPath string) (account cosmosaccount.Account, addr string, err error) {
	isExist, _ := CheckIfAccountExists(accountName, client, addressPrefix, accountPath)
	if isExist {
		registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
		if err != nil {
			fmt.Println(err)
			return cosmosaccount.Account{}, "", err
		}

		account, err := registry.GetByName(accountName)
		if err != nil {
			return cosmosaccount.Account{}, "", err
		}

		addr, err := account.Address(addressPrefix)
		if err != nil {
			fmt.Println("Failed to get the Address:", err)
			return cosmosaccount.Account{}, "", err
		}

		return account, addr, nil
	} else {
		return cosmosaccount.Account{}, "", fmt.Errorf("account not found")
	}
}


func CreateAccount(accountName string, accountPath string) {

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		Log.Error(fmt.Sprintf("Error creating account registry: %v", err))
		return
	}

	account, mnemonic, err := registry.Create(accountName)
	if err != nil {
		Log.Error(fmt.Sprintf("Error creating account: %v", err))
		return
	}

	newCreatedAccount, err := registry.GetByName(accountName)
	if err != nil {
		Log.Error(fmt.Sprintf("Error getting account: %v", err))
		return
	}

	newCreatedAccountAddr, err := newCreatedAccount.Address("air")
	if err != nil {
		Log.Error(fmt.Sprintf("Error getting address: %v", err))
		return
	}

	// create "account.Name".wallet.json file
	type AccountDetails struct {
		Name       string `json:"name"`
		Mnemonic   string `json:"mnemonic"`
		NewAddress string `json:"address"`
	}
	acc := AccountDetails{
		Name:       accountName,
		Mnemonic:   mnemonic,
		NewAddress: newCreatedAccountAddr,
	}
	accountBytes, err := json.Marshal(acc)
	if err != nil {
		Log.Error(fmt.Sprintf("Failed to marshal account details: %s\n", err))
		return
	}
	fileName := fmt.Sprintf("%s/%s.wallet.json", accountPath, accountName)
	// Create and write the file.
	file, err := os.Create(fileName)
	if err != nil {
		Log.Error(fmt.Sprintf("Failed creating file: %s\n", err))
		return
	}
	defer file.Close()
	_, err = file.Write(accountBytes)
	if err != nil {
		Log.Error(fmt.Sprintf("Failed writing to file: %s\n", err))
		return
	}
	Log.Info("File written successfully:" + fileName)

	Log.Info(fmt.Sprintf("Account created: %s", account.Name))
	Log.Info(fmt.Sprintf("Mnemonic: %s", mnemonic))
	Log.Info(fmt.Sprintf("Address: %s", newCreatedAccountAddr))
	Log.Info("Please save this mnemonic key for account recovery")
	Log.Info("Please save this address for future reference")

}
