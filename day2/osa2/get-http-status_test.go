package main

import "testing"

func TestCheckConnection(t *testing.T) {
	ans := checkConnection("https://google.com/")
	if ans != "200 OK" {
		t.Errorf("Check website status!")
	}
}