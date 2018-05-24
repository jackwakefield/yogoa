package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type PositionType int32

const (
	PositionRelative = PositionType(yoga.PositionTypeRelative)
	PositionAbsolute = PositionType(yoga.PositionTypeAbsolute)
)

func (p PositionType) String() string {
	return yoga.PositionTypeToString(yoga.PositionType(p))
}
