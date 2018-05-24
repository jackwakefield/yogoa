package yogoa

import (
	"github.com/jackwakefield/yogoa/yoga"
)

type NodeCloned func(oldNode *Node, newNode *Node, parent *Node, childIndex int32)

type Config struct {
	ref     yoga.ConfigRef
	context interface{}

	setClonedListener bool
	clonedListener    NodeCloned
}

var DefaultConfig = newConfig(yoga.ConfigGetDefault())

func NewConfig() *Config {
	ref := yoga.ConfigNew()
	return newConfig(ref)
}

func newConfig(ref yoga.ConfigRef) *Config {
	return &Config{
		ref: ref,
	}
}

func ConfigCount() int {
	return int(yoga.ConfigGetInstanceCount())
}

func (c *Config) Free() {
	if c.ref != nil {
		yoga.ConfigFree(c.ref)
		c.ref = nil
	}
}

func (c *Config) Copy(dest *Config) {
	if c.ref != nil && dest.ref != nil {
		yoga.ConfigCopy(dest.ref, c.ref)
	}
}

func (c *Config) SetPointScaleFactor(pixelsInPoint float32) {
	if c.ref != nil {
		yoga.ConfigSetPointScaleFactor(c.ref, pixelsInPoint)
	}
}

func (c *Config) SetUseLegacyStretchBehaviour(legacy bool) {
	if c.ref != nil {
		yoga.ConfigSetUseLegacyStretchBehaviour(c.ref, legacy)
	}
}

func (c *Config) SetExperimentEnabled(feature Experiment, enabled bool) {
	if c.ref != nil {
		yoga.ConfigSetExperimentalFeatureEnabled(c.ref, yoga.ExperimentalFeature(feature), enabled)
	}
}

func (c *Config) IsExperimentEnabled(feature Experiment) bool {
	if c.ref != nil {
		return yoga.ConfigIsExperimentalFeatureEnabled(c.ref, yoga.ExperimentalFeature(feature))
	}
	return false
}

func (c *Config) UseWebDefaults() bool {
	if c.ref != nil {
		return yoga.ConfigGetUseWebDefaults(c.ref)
	}
	return false
}

func (c *Config) SetUseWebDefaults(enabled bool) {
	if c.ref != nil {
		yoga.ConfigSetUseWebDefaults(c.ref, enabled)
	}
}

func (c *Config) Context() interface{} {
	return c.context
}

func (c *Config) SetContext(context interface{}) {
	c.context = context
}

func (c *Config) NodeCloned() NodeCloned {
	return c.clonedListener
}

func (c *Config) SetNodeCloned(listener NodeCloned) {
	if c.ref != nil {
		if !c.setClonedListener {
			yoga.ConfigSetNodeClonedFunc(c.ref, c.onCloned)
			c.setClonedListener = true
		}
		c.clonedListener = listener
	}
}

func (c *Config) onCloned(oldRef yoga.NodeRef, newRef yoga.NodeRef, parentRef yoga.NodeRef, childIndex int32) {
	if c.ref != nil && c.clonedListener != nil {
		oldNode := newNode(oldRef)
		node := newNode(newRef)
		parentNode := newNode(parentRef)

		c.clonedListener(oldNode, node, parentNode, childIndex)
	}
}

func (c *Config) Assert(condition bool, message string) {
	if c.ref != nil {
		yoga.AssertWithConfig(c.ref, condition, message)
	}
}
