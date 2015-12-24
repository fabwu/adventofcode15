package main

import (
"testing"
)

func TestCalculateHash1(t *testing.T) {
	expected := 609043
	actual := calculateHash("abcdef", 5)
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCalculateHash2(t *testing.T) {
	expected := 1048970
	actual := calculateHash("pqrstuv", 5)
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}