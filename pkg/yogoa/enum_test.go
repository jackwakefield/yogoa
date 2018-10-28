package yogoa

import (
	"testing"

	"github.com/jackwakefield/yogoa/pkg/yoga"
	"github.com/stretchr/testify/assert"
)

func TestEnumAlign(t *testing.T) {
	values := map[Align]yoga.Align{
		AlignAuto:         yoga.AlignAuto,
		AlignFlexStart:    yoga.AlignFlexStart,
		AlignCenter:       yoga.AlignCenter,
		AlignFlexEnd:      yoga.AlignFlexEnd,
		AlignStretch:      yoga.AlignStretch,
		AlignBaseline:     yoga.AlignBaseline,
		AlignSpaceBetween: yoga.AlignSpaceBetween,
		AlignSpaceAround:  yoga.AlignSpaceAround,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.AlignToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumDimension(t *testing.T) {
	values := map[Dimension]yoga.Dimension{
		DimensionWidth:  yoga.DimensionWidth,
		DimensionHeight: yoga.DimensionHeight,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.DimensionToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumDirection(t *testing.T) {
	values := map[Direction]yoga.Direction{
		DirectionInherit: yoga.DirectionInherit,
		DirectionLTR:     yoga.DirectionLTR,
		DirectionRTL:     yoga.DirectionRTL,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.DirectionToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumDisplay(t *testing.T) {
	values := map[Display]yoga.Display{
		DisplayFlex: yoga.DisplayFlex,
		DisplayNone: yoga.DisplayNone,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.DisplayToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumEdge(t *testing.T) {
	values := map[Edge]yoga.Edge{
		EdgeLeft:       yoga.EdgeLeft,
		EdgeTop:        yoga.EdgeTop,
		EdgeRight:      yoga.EdgeRight,
		EdgeBottom:     yoga.EdgeBottom,
		EdgeStart:      yoga.EdgeStart,
		EdgeEnd:        yoga.EdgeEnd,
		EdgeHorizontal: yoga.EdgeHorizontal,
		EdgeVertical:   yoga.EdgeVertical,
		EdgeAll:        yoga.EdgeAll,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.EdgeToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumExperiment(t *testing.T) {
	values := map[Experiment]yoga.ExperimentalFeature{
		ExperimentWebFlexBasis: yoga.ExperimentalFeatureWebFlexBasis,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.ExperimentalFeatureToString(theirs), ours.String(),
			"string representations should be equal")
	}
}

func TestEnumFlexDirection(t *testing.T) {
	values := map[FlexDirection]yoga.FlexDirection{
		FlexDirectionColumn:        yoga.FlexDirectionColumn,
		FlexDirectionColumnReverse: yoga.FlexDirectionColumnReverse,
		FlexDirectionRow:           yoga.FlexDirectionRow,
		FlexDirectionRowReverse:    yoga.FlexDirectionRowReverse,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.FlexDirectionToString(theirs), ours.String(),
			"string representations should be equal")
	}
}

func TestEnumJustify(t *testing.T) {
	values := map[Justify]yoga.Justify{
		JustifyFlexStart:    yoga.JustifyFlexStart,
		JustifyCenter:       yoga.JustifyCenter,
		JustifyFlexEnd:      yoga.JustifyFlexEnd,
		JustifySpaceBetween: yoga.JustifySpaceBetween,
		JustifySpaceAround:  yoga.JustifySpaceAround,
		JustifySpaceEvenly:  yoga.JustifySpaceEvenly,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.JustifyToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumMeasureMode(t *testing.T) {
	values := map[MeasureMode]yoga.MeasureMode{
		MeasureModeUndefined: yoga.MeasureModeUndefined,
		MeasureModeExactly:   yoga.MeasureModeExactly,
		MeasureModeAtMost:    yoga.MeasureModeAtMost,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.MeasureModeToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumOverflow(t *testing.T) {
	values := map[Overflow]yoga.Overflow{
		OverflowVisible: yoga.OverflowVisible,
		OverflowHidden:  yoga.OverflowHidden,
		OverflowScroll:  yoga.OverflowScroll,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.OverflowToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumPositionType(t *testing.T) {
	values := map[PositionType]yoga.PositionType{
		PositionRelative: yoga.PositionTypeRelative,
		PositionAbsolute: yoga.PositionTypeAbsolute,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.PositionTypeToString(theirs), ours.String(),
			"string representations should be equal")
	}
}

func TestEnumPrintOption(t *testing.T) {
	values := map[PrintOption]yoga.PrintOptions{
		PrintOptionLayout:   yoga.PrintOptionsLayout,
		PrintOptionStyle:    yoga.PrintOptionsStyle,
		PrintOptionChildren: yoga.PrintOptionsChildren,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.PrintOptionsToString(theirs), ours.String(),
			"string representations should be equal")
	}
}

func TestEnumUnit(t *testing.T) {
	values := map[Unit]yoga.Unit{
		UnitUndefined: yoga.UnitUndefined,
		UnitPoint:     yoga.UnitPoint,
		UnitPercent:   yoga.UnitPercent,
		UnitAuto:      yoga.UnitAuto,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.UnitToString(theirs), ours.String(), "string representations should be equal")
	}
}

func TestEnumWrap(t *testing.T) {
	values := map[Wrap]yoga.Wrap{
		WrapNone:    yoga.WrapNoWrap,
		WrapWrap:    yoga.WrapWrap,
		WrapReverse: yoga.WrapWrapReverse,
	}
	for ours, theirs := range values {
		assert.EqualValues(t, theirs, ours, "values should be equal")
		assert.EqualValues(t, yoga.WrapToString(theirs), ours.String(), "string representations should be equal")
	}
}
