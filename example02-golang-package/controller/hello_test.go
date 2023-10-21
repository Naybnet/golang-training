package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("appleboy")
	expected := "Hi, appleboy"
	if hello != expected {
		t.Errorf("Testing fail <%s> should be <%s>", hello, expected)
	}

	hello = HelloWorld("appleboy ")
	if hello != expected {
		t.Errorf("Testing fail <%s> should be <%s>", hello, expected)
	}
}
