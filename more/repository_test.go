package more_test

import (
	"more"
	"testing"
)

func TestMysqlRepository(t *testing.T) {
	repo := more.NewMysqlRepository()

	// Test Insert
	if !repo.Insert("item1") {
		t.Errorf("MysqlRepository Insert failed")
	}
	if !repo.Items["item1"] {
		t.Errorf("MysqlRepository Insert did not store the item correctly")
	}

	// Test Delete
	if !repo.Delete("item1") {
		t.Errorf("MysqlRepository Delete failed")
	}
	if repo.Items["item1"] {
		t.Errorf("MysqlRepository Delete did not remove the item correctly")
	}
}

func TestPostgresRepository(t *testing.T) {
	repo := more.NewPostgresRepository()

	// Test Insert
	if !repo.Insert("item1") {
		t.Errorf("PostgresRepository Insert failed")
	}
	if !repo.Items["item1"] {
		t.Errorf("PostgresRepository Insert did not store the item correctly")
	}

	// Test Delete
	if !repo.Delete("item1") {
		t.Errorf("PostgresRepository Delete failed")
	}
	if repo.Items["item1"] {
		t.Errorf("PostgresRepository Delete did not remove the item correctly")
	}
}
