package yogoa

import "github.com/jackwakefield/yogoa/pkg/yoga"

type Align int32

const (
	AlignAuto         = Align(yoga.AlignAuto)
	AlignFlexStart    = Align(yoga.AlignFlexStart)
	AlignCenter       = Align(yoga.AlignCenter)
	AlignFlexEnd      = Align(yoga.AlignFlexEnd)
	AlignStretch      = Align(yoga.AlignStretch)
	AlignBaseline     = Align(yoga.AlignBaseline)
	AlignSpaceBetween = Align(yoga.AlignSpaceBetween)
	AlignSpaceAround  = Align(yoga.AlignSpaceAround)
)

func (a Align) String() string {
	return yoga.AlignToString(yoga.Align(a))
}
