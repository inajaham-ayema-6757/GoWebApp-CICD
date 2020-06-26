package main

import (
	"testing"
)

func Test1_greeting(t *testing.T) {
	greet := greeting()
	if greet != "Howdy Folks..xD!!!" {
		t.Errorf("Unexpected Output from greeting function, got: %s, want: %s.", greet, "Howdy Folks..xD!!!")
	}
}

func Test2_introduction(t *testing.T) {
	intro_msg := greeting()
	if intro_msg != "My name is Ameya Makarand Mahajani" {
		t.Errorf("Unexpected Output from introduction function, got: %s, want: %s.", intro_msg, "My name is Ameya Makarand Mahajani")
	}
}
