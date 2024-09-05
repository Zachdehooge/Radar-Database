package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	input := strings.NewReader("jane")
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	name := scanner.Text()
	if len(name) == 0 {
		t.Log("empty input")
	}
	t.Logf("You entered: %s\n", name)
}
func TestHello(t *testing.T) {
	got := "Hello, world"
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
