package checkproxy

import (
	"strconv"
	"testing"
	"time"
)

func TestCheckSocksProxy(t *testing.T) {

	nonProxy := CheckSocksProxy("0.0.0.0:0000", time.Duration(5*time.Second))

	// Return false for non proxy
	if nonProxy {
		t.Errorf("Expected false, got: %s", strconv.FormatBool(nonProxy))
	}
}
