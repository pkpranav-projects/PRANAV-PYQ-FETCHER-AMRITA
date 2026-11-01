package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRoot(t *testing.T) {

	// Test RootCmd
	if _, ok := interface{}(Cmd).(*cobra.Command); !ok {
		t.Errorf("RootCmd is not of type *cobra.Command")
	}

	if Cmd.Use != "pranav-pyq" {
		t.Errorf("Expected Use: 'pranav-pyq', Got: %v", Cmd.Use)
	}

	if Cmd.Short != "PRANAV PYQ FETCHER AMRITA" {
		t.Errorf("Expected Short: 'PRANAV PYQ FETCHER AMRITA', Got: %v", Cmd.Short)
	}
}
