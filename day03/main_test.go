package main

import (
"testing"
)

func TestSantaDeliver1(t *testing.T) {
	houses := santaDeliver(">")
	if houses != 2 {
		t.Error("Expected 2, got ", houses)
	}
}

func TestSantaDeliver2(t *testing.T) {
	houses := santaDeliver("^>v<")
	if houses != 4 {
		t.Error("Expected 4, got ", houses)
	}
}

func TestSantaDeliver3(t *testing.T) {
	houses := santaDeliver("^v^v^v^v^v")
	if houses != 2 {
		t.Error("Expected 2, got ", houses)
	}
}

func TestRoboSantaDeliver1(t *testing.T) {
	houses := roboSantaDeliver("^v")
	if houses != 3 {
		t.Error("Expected 3, got ", houses)
	}
}

func TestRoboSantaDeliver2(t *testing.T) {
	houses := roboSantaDeliver("^>v<")
	if houses != 3 {
		t.Error("Expected 3, got ", houses)
	}
}

func TestRoboSantaDeliver3(t *testing.T) {
	houses := roboSantaDeliver("^v^v^v^v^v")
	if houses != 11 {
		t.Error("Expected 11, got ", houses)
	}
}

