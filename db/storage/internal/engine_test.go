package internal

import (
	"testing"
)

func TestEngine_SetAndGet(t *testing.T) {
	engine := NewEngine()

	key := "testKey"
	value := "testValue"

	engine.Set(key, value)

	retrievedValue, err := engine.Get(key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if retrievedValue != value {
		t.Errorf("Expected value %s, received %s", value, retrievedValue)
	}
}

func TestEngine_Get_NonExistentKey(t *testing.T) {
	engine := NewEngine()

	_, err := engine.Get("nonExistentKey")
	if err == nil {
		t.Fatalf("Error expected when retrieving a non-existent key.")
	}
}

func TestEngine_Del(t *testing.T) {
	engine := NewEngine()

	key := "testKey"
	value := "testValue"

	engine.Set(key, value)
	engine.Del(key)

	_, err := engine.Get(key)
	if err == nil {
		t.Fatalf("Error expected after deleting key")
	}
}
