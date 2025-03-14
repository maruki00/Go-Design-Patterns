package more_test

import (
	"_/home/user/dev/Projects/Go/public/design-patterns/more"
	"testing"
)

func TestMysqlRepository(t *testing.T) {
	repo := more.NewMysqlRepository()

	// Test Insert
	if !repo.Insert("item1") {
		t.Errorf("MysqlRepository Insert failed")
	}
	if !repo.items["item1"] {
		t.Errorf("MysqlRepository Insert did not store the item correctly")
	}

	// Test Delete
	if !repo.Delete("item1") {
		t.Errorf("MysqlRepository Delete failed")
	}
	if repo.items["item1"] {
		t.Errorf("MysqlRepository Delete did not remove the item correctly")
	}
}

func TestPostgresRepository(t *testing.T) {
	repo := NewPostgresRepository()

	// Test Insert
	if !repo.Insert("item1") {
		t.Errorf("PostgresRepository Insert failed")
	}
	if !repo.items["item1"] {
		t.Errorf("PostgresRepository Insert did not store the item correctly")
	}

	// Test Delete
	if !repo.Delete("item1") {
		t.Errorf("PostgresRepository Delete failed")
	}
	if repo.items["item1"] {
		t.Errorf("PostgresRepository Delete did not remove the item correctly")
	}
}

func TestService(t *testing.T) {
	repo := NewMysqlRepository()
	service := NewService(repo)

	// Test Create via Service
	if !service.Create("item1") {
		t.Errorf("Service Create failed")
	}
	if !repo.items["item1"] {
		t.Errorf("Service Create did not store the item correctly in the repository")
	}
}
