package e20

import (
	"log"
	"testing"
)

func setup() {
	log.Println("setup")
}

func tearDown() {
	log.Println("tearDown")
}

func TestFactorialOfPositive(t *testing.T) {
	setup()
	defer tearDown()
	if expected, actual := 120, Factorial(5); expected != actual {
		t.Fatalf("Expecting %v, got %v\n", expected, actual)
	}
}

func TestFactorialOfNegative(t *testing.T) {
	setup()
	defer tearDown()
	if expected, actual := 1, Factorial(-5); expected != actual {
		t.Fatalf("Expecting %v, got %v\n", expected, actual)
	}
}
