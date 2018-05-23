package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type FlexDirection int32

const (
	FlexDirectionColumn        = FlexDirection(yoga.FlexDirectionColumn)
	FlexDirectionColumnReverse = FlexDirection(yoga.FlexDirectionColumnReverse)
	FlexDirectionRow           = FlexDirection(yoga.FlexDirectionRow)
	FlexDirectionRowReverse    = FlexDirection(yoga.FlexDirectionRowReverse)
)

func (d FlexDirection) String() string {
	return yoga.FlexDirectionToString(yoga.FlexDirection(d))
}
