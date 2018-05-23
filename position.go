package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type PositionType int32

const (
	PositionTypeRelative = PositionType(yoga.PositionTypeRelative)
	PositionTypeAbsolute = PositionType(yoga.PositionTypeAbsolute)
)

func (p PositionType) String() string {
	return yoga.PositionTypeToString(yoga.PositionType(p))
}
