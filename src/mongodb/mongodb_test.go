package mongodb

import (
	"os"
	"testing"

	"pranavt61/icon-graphql-api/graph/model"
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
	block, err := DBConnection.FindBlockByHash("TestHash1")
	if err != nil {
		t.Errorf("Failed to query mongodb: %s", err.Error())

		// Fail
		return
	}

	if block.Hash != "TestHash1" {
		t.Errorf("Failed to validate mongodb data")
	}

	// Pass
	return
}
