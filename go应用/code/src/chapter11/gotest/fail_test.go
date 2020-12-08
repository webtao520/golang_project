package gotest

import (
	"fmt"
	"testing"
)

func TestFailNow(t *testing.T) {
	t.FailNow()
}

func TestFail(t *testing.T) {

	fmt.Println("before fail")

	t.Fail()

	fmt.Println("after fail")
}
