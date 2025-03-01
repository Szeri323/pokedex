package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	intrval := 5 * time.Second
	cache := NewCache(intrval)

	if cache.interval != intrval {
		t.Errorf("expected interval %v, got %v", intrval, cache.interval)
	}

	if cache.cache == nil {
		t.Errorf("expected cache to be initialized, but got nil")
	}
}

func TestAdd(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)

	testKey := "https://test.com"
	testValue := []byte("AAA")

	cache.Add(testKey, testValue)

	retrivedValue, exists := cache.Get(testKey)

	if !exists {
		t.Errorf("Expected key %s to exist in cache but it doesn't", testKey)
	}
	if string(retrivedValue) != string(testValue) {
		t.Errorf("Expected value %s, got %s", string(testValue), string(retrivedValue))
	}
}

func TestReapLoop(t *testing.T) {

}


/*
func TestReapLoop(t *testing.T) {
    // Use a very short interval for testing to avoid long test duration
    const interval = 5 * time.Millisecond
    const waitTime = interval + 10*time.Millisecond // Wait a bit longer than the interval
    
    cache := NewCache(interval)
    testKey := "https://example.com"
    testValue := []byte("testdata")
    
    // Add an item to the cache
    cache.Add(testKey, testValue)
    
    // Verify it was added
    val, exists := cache.Get(testKey)
    if !exists {
        t.Errorf("expected to find key immediately after adding")
        return
    }
    if string(val) != string(testValue) {
        t.Errorf("expected value %s, got %s", string(testValue), string(val))
        return
    }
    
    // Wait for longer than the interval
    time.Sleep(waitTime)
    
    // Verify the item was removed by the reap loop
    _, exists = cache.Get(testKey)
    if exists {
        t.Errorf("expected key to be removed after interval")
        return
    }
}
	This test:

Creates a cache with a very short interval (5 milliseconds)
Adds an item to the cache
Verifies the item exists in the cache
Waits longer than the interval (15 milliseconds)
Verifies the item has been removed by the reap loop
The key here is using a very short interval and waiting just long enough for the reap loop to run, but not so long that the test becomes slow.
*/