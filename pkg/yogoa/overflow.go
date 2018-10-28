package yogoa

import "github.com/jackwakefield/yogoa/pkg/yoga"

type Overflow int32

const (
	OverflowVisible = Overflow(yoga.OverflowVisible)
	OverflowHidden  = Overflow(yoga.OverflowHidden)
	OverflowScroll  = Overflow(yoga.OverflowScroll)
)

func (o Overflow) String() string {
	return yoga.OverflowToString(yoga.Overflow(o))
}
