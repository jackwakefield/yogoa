package yogoa

import "testing"

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	config.Free()
}
