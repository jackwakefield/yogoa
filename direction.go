package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type Direction int32

const (
	DirectionInherit = Direction(yoga.DirectionInherit)
	DirectionLTR     = Direction(yoga.DirectionLTR)
	DirectionRTL     = Direction(yoga.DirectionRTL)
)

func (d Direction) String() string {
	return yoga.DirectionToString(yoga.Direction(d))
}
