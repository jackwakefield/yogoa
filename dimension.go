package yogoa

import (
	"github.com/jackwakefield/yogoa/yoga"
)

type Dimension int32

const (
	DimensionWidth  = Dimension(yoga.DimensionWidth)
	DimensionHeight = Dimension(yoga.DimensionHeight)
)

func (d Dimension) String() string {
	return yoga.DimensionToString(yoga.Dimension(d))
}
