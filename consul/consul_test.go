package consul

import (
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	err := Register("metrics-hub")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(31 * time.Second)
}
