package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	var buf bytes.Buffer
	versionCmd.SetOut(&buf)
	versionCmd.SetErr(&buf)
	versionCmd.Run(versionCmd, nil)
	if !strings.Contains(buf.String(), "PRANAV PYQ FETCHER AMRITA v2.0.0") {
		t.Errorf("got %q, want it to contain PRANAV PYQ FETCHER AMRITA v2.0.0", buf.String())
	}
}
