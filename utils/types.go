package utils

// Config holds configuration details for the blockchain connection and accounts.
type Config struct {
	AccountPath      string
	JsonRPC          string
	Token            string
	ChainPrefix      string
	AccountNames     []string
	AccountAddresses []string
}
