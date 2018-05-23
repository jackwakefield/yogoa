package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type MeasureMode int32

const (
	MeasureModeUndefined = MeasureMode(yoga.MeasureModeUndefined)
	MeasureModeExactly   = MeasureMode(yoga.MeasureModeExactly)
	MeasureModeAtMost    = MeasureMode(yoga.MeasureModeAtMost)
)

func (m MeasureMode) String() string {
	return yoga.MeasureModeToString(yoga.MeasureMode(m))
}

func CanUseCachedMeasurement(
	widthMode MeasureMode, width float32,
	heightMode MeasureMode, height float32,
	lastWidthMode MeasureMode, lastWidth float32,
	lastHeightMode MeasureMode, lastHeight float32,
	lastComputedWidth float32, lastComputedHeight float32,
	marginRow float32, marginColumn float32,
	config *Config,
) bool {
	if config != nil && config.ref != nil {
		return yoga.NodeCanUseCachedMeasurement(
			yoga.MeasureMode(widthMode), width,
			yoga.MeasureMode(heightMode), height,
			yoga.MeasureMode(lastWidthMode), lastWidth,
			yoga.MeasureMode(lastHeightMode), lastHeight,
			lastComputedWidth, lastComputedHeight,
			marginRow, marginColumn,
			config.ref,
		)
	}
	return false
}
