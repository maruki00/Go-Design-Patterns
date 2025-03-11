package creational

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	obj1 := GetObject()
	obj2 := GetObject()

	if obj1 != obj2 {
		t.Error("Expected both objects to be the same instance")
	}

	obj1.SetId(100)

	if obj2.GetId() != 100 {
		t.Errorf("Expected id to be 100, but got %d", obj2.GetId())
	}
}
