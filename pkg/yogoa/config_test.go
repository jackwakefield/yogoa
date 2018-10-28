package yogoa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert.Equal(t, 1, ConfigCount(), "ConfigCount() should be 1")

	config := NewConfig()
	assert.NotNil(t, config, "config should not be nil")
	assert.NotNil(t, config.ref, "config.ref should not be nil")
	assert.Equal(t, 2, ConfigCount(), "ConfigCount() should be 2")

	config.Free()
	assert.Nil(t, config.ref, "config.ref should be nil")
	assert.Equal(t, 1, ConfigCount(), "ConfigCount() should be 1")
}

func TestConfigFree(t *testing.T) {
	config := NewConfig()
	assert.NotNil(t, config.ref, "config.ref should not be nil")
	config.Free()
	assert.Nil(t, config.ref, "config.ref should be nil")

	assert.NotPanics(t, func() {
		config.Free()
	}, "config.Free() should not panic when called multiple times")
}

func TestConfigCopy(t *testing.T) {
	a := NewConfig()
	assert.Equal(t, false, a.UseWebDefaults(), "a.UseWebDefaults() should be false")
	a.SetUseWebDefaults(true)
	assert.Equal(t, true, a.UseWebDefaults(), "a.UseWebDefaults() should be true")

	b := NewConfig()
	assert.Equal(t, false, b.UseWebDefaults(), "b.UseWebDefaults() should be false")
	a.Copy(b)
	assert.Equal(t, true, b.UseWebDefaults(), "b.UseWebDefaults() should be true")

	b.Free()
	assert.NotPanics(t, func() {
		a.Copy(b)
	}, "a.Copy(b) should not panic when copied after b has been freed")

	a.Free()
	assert.NotPanics(t, func() {
		a.Copy(b)
	}, "a.Copy(b) should not panic when copied after both a and b have been freed")
}

func TestConfigSetPointScaleFactor(t *testing.T) {
	config := NewConfig()
	config.SetPointScaleFactor(1)
	config.Free()

	assert.NotPanics(t, func() {
		config.SetPointScaleFactor(1)
	}, "config.SetPointScaleFactor() should not panic after free")
}

func TestConfigSetUseLegacyStretchBehaviour(t *testing.T) {
	config := NewConfig()
	config.SetUseLegacyStretchBehaviour(true)
	config.Free()

	assert.NotPanics(t, func() {
		config.SetUseLegacyStretchBehaviour(true)
	}, "config.SetUseLegacyStretchBehaviour() should not panic after free")
}

func TestConfigExperimentEnabled(t *testing.T) {
	config := NewConfig()
	assert.Equal(t, false, config.IsExperimentEnabled(ExperimentWebFlexBasis),
		"config.IsExperimentEnabled(ExperimentWebFlexBasis) should equal false")

	config.SetExperimentEnabled(ExperimentWebFlexBasis, true)
	assert.Equal(t, true, config.IsExperimentEnabled(ExperimentWebFlexBasis),
		"config.IsExperimentEnabled(ExperimentWebFlexBasis) should equal true")

	config.Free()

	assert.NotPanics(t, func() {
		config.SetExperimentEnabled(ExperimentWebFlexBasis, true)
	}, "config.SetExperimentEnabled(ExperimentWebFlexBasis) should not panic after free")

	assert.Equal(t, false, config.IsExperimentEnabled(ExperimentWebFlexBasis),
		"config.IsExperimentEnabled(ExperimentWebFlexBasis) should equal false")
}

func TestConfigContext(t *testing.T) {
	ctx := &struct {
		Foo int32
	}{
		Foo: 0x1234ABCD,
	}

	config := NewConfig()
	config.SetContext(ctx)
	assert.Equal(t, ctx, config.Context(), "config.Context() should equal ctx")
	config.SetContext(nil)
	assert.Nil(t, config.Context(), "config.Context() should equal nil")
	config.Free()

	assert.NotPanics(t, func() {
		config.SetContext(ctx)
	}, "config.SetContext() should not panic after free")

	assert.Nil(t, config.Context(), "config.Context() should equal nil")
}
