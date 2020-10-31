package sysop_test

import (
	"testing"

	"bakku.dev/dotf/sysop"
)

func TestPathExists(t *testing.T) {
	op := sysop.Provider{}

	if op.PathExists("sysop.go") == false {
		t.Fatal("expected sysop.go to exist")
	}

	if op.PathExists("invalid.invalid") {
		t.Fatal("expected invalid.invalid not to exist")
	}

	if op.PathExists("../sysop") == false {
		t.Fatal("expected ../sysop to exist")
	}
}
