package main

import (
        "testing"
)

func TestHello(t *testing.T) {
        expectedStr := "From sandbox!"
        result := hello()
        if result != expectedStr {
                t.Fatalf("Expected %s, got %s", expectedStr, result)
        }
}
