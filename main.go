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
			"account3",
			"account4",
		},
		AccountAddresses: []string{
			"air1mue344jmwztfgpjsf2v34utwtgpjqv00yzv5kz",
			"air1xgj29fgn22wuy0hpm8jy64k3z5setpwtdmec77",
			"air1w2c7dujdatxvupp7qy4x5wcsk295alher9059g",
			"air1w3w9eq0l47cx4xk8erz3hz4zstpyfhql06532y",
		},
	}

	setupAccounts(config)
	checkBalances(config)
	// accountsData := connectAllClients(config)


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
