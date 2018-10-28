package yogoa

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"

	"github.com/jackwakefield/yogoa/pkg/yoga"
)

type CloneNode func(oldNode *Node, parent *Node, childIndex int32) *Node

type Config struct {
	ref yoga.ConfigRef
	ctx interface{}

	onCloneNode CloneNode
}

var configsByContext sync.Map

var DefaultConfig = newConfig(yoga.ConfigGetDefault())

func NewConfig() *Config {
	return newConfig(yoga.ConfigNew())
}

func configFromRef(ref yoga.ConfigRef) *Config {
	if ptr := yoga.ConfigGetContext(ref); ptr != nil {
		if value, ok := configsByContext.Load(ptr); ok {
			if config, ok := value.(*Config); ok {
				return config
			}
		}
	}
	return nil
}

func newConfig(ref yoga.ConfigRef) *Config {
	config := configFromRef(ref)
	if config == nil {
		config = &Config{ref: ref}

		var ctx unsafe.Pointer = C.malloc(C.size_t(1))
		yoga.ConfigSetContext(ref, ctx)
		configsByContext.Store(ctx, config)
	}
	return config
}

func ConfigCount() int {
	return int(yoga.ConfigGetInstanceCount())
}

func (c *Config) Free() {
	if c.ref != nil {
		if ptr := yoga.ConfigGetContext(c.ref); ptr != nil {
			if _, ok := configsByContext.Load(ptr); ok {
				configsByContext.Delete(ptr)
				C.free(ptr)
			}
		}

		yoga.ConfigFree(c.ref)
		c.ref = nil
	}
}

func (c *Config) Copy(dest *Config) {
	if c.ref != nil && dest.ref != nil {
		yoga.ConfigCopy(dest.ref, c.ref)
	}
}

func (c *Config) Context() interface{} {
	if c.ref != nil {
		return c.ctx
	}
	return nil
}

func (c *Config) SetContext(ctx interface{}) {
	if c.ref != nil {
		c.ctx = ctx
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

func (c *Config) SetCloneNode(listener CloneNode) {
	if c.ref != nil {
		c.onCloneNode = listener

		if listener != nil {
			yoga.ConfigSetCloneNodeFunc(c.ref, onCloneNode)
		}
	}
}

func onCloneNode(configRef yoga.ConfigRef, oldRef yoga.NodeRef, parentRef yoga.NodeRef, childIndex int32) yoga.NodeRef {
	config := configFromRef(configRef)
	if config != nil && config.onCloneNode != nil {
		oldNode := newNode(oldRef)
		parentNode := newNode(parentRef)

		if newNode := config.onCloneNode(oldNode, parentNode, childIndex); newNode != nil {
			return newNode.ref
		}
	}
	return nil
}
