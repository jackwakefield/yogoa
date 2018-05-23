package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type Justify int32

const (
	JustifyFlexStart    = Justify(yoga.JustifyFlexStart)
	JustifyCenter       = Justify(yoga.JustifyCenter)
	JustifyFlexEnd      = Justify(yoga.JustifyFlexEnd)
	JustifySpaceBetween = Justify(yoga.JustifySpaceBetween)
	JustifySpaceAround  = Justify(yoga.JustifySpaceAround)
	JustifySpaceEvenly  = Justify(yoga.JustifySpaceEvenly)
)

func (j Justify) String() string {
	return yoga.JustifyToString(yoga.Justify(j))
}
