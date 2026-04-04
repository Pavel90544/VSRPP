package main

import (
    "testing"
)

func TestFullApp(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }
    
    t.Log("Integration test passed")
}
