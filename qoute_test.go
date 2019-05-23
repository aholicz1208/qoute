package qoute

import (
    "testing"
)

func Test_Say(t *testing.T) {
    want := "hello"

    if got := Say(); got != want {
        t.Errorf("Hello() = %s , want = %s", got, want)
    }
}
