package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, error := IsEmail("Hello")

	if error == nil {
		t.Error("hello is not an email")
	}

	_, error = IsEmail("Derek@aol.com")

	if error == nil {
		t.Error("derek@aol.com is an email")
	}

	_, error = IsEmail("Derek@aol")

	if error == nil {
		t.Error("derek@aol is an email")
	}
}
