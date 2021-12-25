package helper

import "testing"

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("indonesia juara")

	if result != "hello indonesia juara" {
		panic("result is not hello indonesia juara")
	}
}

func TestHelloWorldHafet(t *testing.T)  {
	result := HelloWorld("hafid juara")

	if result != "hello hafid juara" {
		panic("result is not hello hafid juara")
	}
}
