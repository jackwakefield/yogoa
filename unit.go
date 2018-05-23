package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type Unit int32

const (
	UnitUndefined = Unit(yoga.UnitUndefined)
	UnitPoint     = Unit(yoga.UnitPoint)
	UnitPercent   = Unit(yoga.UnitPercent)
	UnitAuto      = Unit(yoga.UnitAuto)
)

func (u Unit) String() string {
	return yoga.UnitToString(yoga.Unit(u))
}
