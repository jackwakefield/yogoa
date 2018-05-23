package yogoa

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode()
	node.Free()
}
