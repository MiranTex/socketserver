package database_test

import (
	"socket_project/cluster_module/database"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {

	db := database.GetSqliteConnection()

	if db == nil {
		t.Errorf("Expected db connection to be not nil, got nil")
	}

}
