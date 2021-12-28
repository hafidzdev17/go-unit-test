package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("indonesia juara")

	if result != "hello indonesia juara" {
		t.Error("result must be 'hello indonesia juara' ")
	}
}

func TestHelloWorldHafet(t *testing.T) {
	result := HelloWorld("hafid juara")

	if result != "hello hafid juara" {
		t.Fatal("result must be 'hello hafid juara' ")
	}
}
