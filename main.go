package main

import (
	"fmt"
	"github.com/saatvik333/junction-go-client-template/junction"
	"github.com/saatvik333/junction-go-client-template/utils"
	"log"
)

func main() {
	// Define configuration
	config := utils.Config{
		AccountPath: "./accounts",
		JsonRPC:     "http://localhost:26657",
		Token:       "stake",
		ChainPrefix: "air",
		AccountNames: []string{
			"account1",
			"account2",
			
		},
		AccountAddresses: []string{
			"air1cxfn6mrqz98h0rkq8pkxg47jf33e5jhh093pjq",
			"air160kc8z27qhx8ef4aagd0usf63h5mz952qklfen",
			
		},
	}

	setupAccounts(config)
	checkBalances(config)
	 //accountsData := connectAllClients(config)


}

// setupAccounts creates accounts using the provided configuration.
func setupAccounts(cfg utils.Config) {
	for _, name := range cfg.AccountNames {
		utils.CreateAccount(name, cfg.AccountPath)
		log.Printf("Failed to create account %s:", name)

	}
}

// checkBalances retrieves and logs the balance of each account.
func checkBalances(cfg utils.Config) {
	for _, addr := range cfg.AccountAddresses {
		success, balance, err := utils.CheckBalance(cfg.JsonRPC, cfg.Token, addr)
		if err != nil {
			log.Printf("Error checking balance for %s: %v", addr, err)
			continue
		}
		if !success {
			log.Printf("Failed to retrieve balance for %s", addr)
			continue
		}
		utils.Log.Info(fmt.Sprintf("%s balance: %d", addr, balance))
	}
}

// connectAllClients connects to all accounts and returns their connection data.
func connectAllClients(cfg utils.Config) []utils.AccountData {
	var accountsData []utils.AccountData
	for _, name := range cfg.AccountNames {
		account, addr, client := junction.ClientConnect(cfg.AccountPath, name, cfg.ChainPrefix, cfg.JsonRPC)
		accountsData = append(accountsData, utils.AccountData{
			Account: account,
			Addr:    addr,
			Client:  client,
		})
	}
	return accountsData
}
