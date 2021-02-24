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
		return
	}

	// Pass
	return
}
