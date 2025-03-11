package creational_test

import (
	"creational"
	"testing"
)

func TestQuery(t *testing.T) {

	q := &creational.Query{}
	q.Table("tab1").Select([]string{"id", "name"}).Where("id = 1").Limit(10).Offset(10).Get()
}
