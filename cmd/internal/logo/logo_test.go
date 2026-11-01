package logo

import (
	"strings"
	"testing"
)

func TestLogo(t *testing.T) {
	t.Run("TestLogoASCII", func(t *testing.T) {
		t.Parallel()
		if !strings.Contains(LOGO_ASCII, "AMRITA UNIVERSITY") {
			t.Errorf("Expected logo to contain AMRITA UNIVERSITY, got %s", LOGO_ASCII)
		}
	})
}
