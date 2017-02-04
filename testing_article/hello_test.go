package main

import (
  "testing"
)

func TestHello(t *testing.T)  {
  expectedStr := "Hello, Testing!"
  result, _ := hello()
  if result != expectedStr {
    t.Fatalf("Expected %s, got %s", expectedStr, result)
  }
}
