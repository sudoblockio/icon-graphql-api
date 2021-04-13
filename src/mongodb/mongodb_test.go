package mongodb

import (
	"os"
	"testing"
)

func TestConnectClient(t *testing.T) {

	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Pass
	return
}

func TestFindBlockByHash(t *testing.T) {
	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Find block
	// Mock data inserted at mongodb/init/mock.js
	block, err := DBConnection.FindBlockByHash("4d7b862e4f49d23ca445a5df4e44585c9822563f115e60d0f544372797192102")
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if block.Hash != "4d7b862e4f49d23ca445a5df4e44585c9822563f115e60d0f544372797192102" ||
		block.Number != 10000019 {
		t.Errorf("Failed to validate mongodb data")

		// Fail
		return
	}

	// Pass
	return
}

func TestBlocks(t *testing.T) {
	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Find block
	// Mock data inserted at mongodb/init/mock.js
	blocks, err := DBConnection.Blocks(0, 1)
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if len(blocks) == 0 {
		t.Errorf("Failed to retrieve any blocks")

		// Fail
		return
	}

	// Pass
	return
}

func TestFindTransactionByHash(t *testing.T) {
	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Find block
	// Mock data inserted at mongodb/init/mock.js
	transaction, err := DBConnection.FindTransactionByHash("0xc3528666aa2f5e5cca5fe9c8121ccf1593eb844bf56246ec1a9d3ec202bd3f77")
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if transaction.Hash != "0xc3528666aa2f5e5cca5fe9c8121ccf1593eb844bf56246ec1a9d3ec202bd3f77" {
		t.Errorf("Failed to validate transaction")

		// Fail
		return
	}

	// Pass
	return

}

func TestAllTransactions(t *testing.T) {
	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Find block
	// Mock data inserted at mongodb/init/mock.js
	transactions, err := DBConnection.AllTransactions()
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if len(transactions) == 0 {
		t.Errorf("Failed to retrieve any transactions")

		// Fail
		return
	}

	// Pass
	return
}

func TestFindLogsByTransactionHash(t *testing.T) {
	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")

	err := ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)
	if err != nil {
		t.Errorf("Failed to connect mongo client: %s", err.Error())

		// Fail
		return
	}

	// Find block
	// Mock data inserted at mongodb/init/mock.js
	logs, err := DBConnection.FindLogsByTransactionHash("0xc92f6cea710f2e420aab9c44f16c69502f5729f08412c1333ed54979a5ee4200")
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if len(logs) == 0 {
		t.Errorf("Failed to retrieve any logs")

		// Fail
		return
	}

	// Pass
	return
}
