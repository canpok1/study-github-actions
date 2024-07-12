package main

import "testing"

func Test_message(t *testing.T) {
	got := message()
	want := "run sample-app"
	if got != want {
		t.Errorf("message() = %v, want %v", got, want)
	}
}
