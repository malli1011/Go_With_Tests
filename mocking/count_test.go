package mocking

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buff := &bytes.Buffer{}
	CountDown(buff)

	got := buff.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
