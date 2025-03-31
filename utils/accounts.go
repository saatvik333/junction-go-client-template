package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	// "go.uber.org/zap"

	
)

// AccountData stores account connection data.
type AccountData struct {
	Account cosmosaccount.Account
	Addr    string
	Client  cosmosclient.Client
}


// Helper function to log errors consistently
func logError(msg string, err error) {
	// ensureLogger()
	fmt.Sprintf("%s: %v", msg, err)
}

// CheckIfAccountExists verifies if an account exists in the registry
func CheckIfAccountExists(accountName, accountPath, addressPrefix string) (bool, string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		logError("Failed to create account registry", err)
		return false, ""
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		return false, ""
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		logError("Failed to retrieve account address", err)
		return false, ""
	}

	return true, addr
}

// FetchAccount retrieves an existing account by name
func FetchAccount(accountName, accountPath, addressPrefix string) (cosmosaccount.Account, string, error) {
	exists, _ := CheckIfAccountExists(accountName, accountPath, addressPrefix)
	if !exists {
		return cosmosaccount.Account{}, "", fmt.Errorf("account not found")
	}

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		logError("Failed to create account registry", err)
		return cosmosaccount.Account{}, "", err
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		logError("Failed to retrieve account", err)
		return cosmosaccount.Account{}, "", err
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		logError("Failed to retrieve account address", err)
		return cosmosaccount.Account{}, "", err
	}

	return account, addr, nil
}

// CreateAccount generates a new Cosmos account and saves its details to a file
func CreateAccount(accountName, accountPath string) {
	// ensureLogger()

	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		logError("Error creating account registry", err)
		return
	}

	account, mnemonic, err := registry.Create(accountName)
	if err != nil {
		logError("Error creating account", err)
		return
	}

	newAccount, err := registry.GetByName(accountName)
	if err != nil {
		logError("Error retrieving account", err)
		return
	}

	newAccountAddr, err := newAccount.Address("air")
	if err != nil {
		logError("Error retrieving account address", err)
		return
	}

	accountDetails := struct {
		Name     string `json:"name"`
		Mnemonic string `json:"mnemonic"`
		Address  string `json:"address"`
	}{
		Name:     accountName,
		Mnemonic: mnemonic,
		Address:  newAccountAddr,
	}

	accountBytes, err := json.Marshal(accountDetails)
	if err != nil {
		logError("Failed to marshal account details", err)
		return
	}

	fileName := fmt.Sprintf("%s/%s.wallet.json", accountPath, accountName)
	file, err := os.Create(fileName)
	if err != nil {
		logError("Failed to create wallet file", err)
		return
	}
	defer file.Close()

	if _, err = file.Write(accountBytes); err != nil {
		logError("Failed writing to wallet file", err)
		return
	}

	fmt.Printf("Account created successfully: %s", account.Name)
	fmt.Printf("Mnemonic: %s", mnemonic)
	fmt.Printf("Address: %s", newAccountAddr)
	fmt.Printf("Wallet file saved at: %s", fileName)
	fmt.Println("Save this mnemonic key securely for account recovery.")
}
