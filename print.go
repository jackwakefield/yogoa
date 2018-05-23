package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type PrintOptions int32

const (
	PrintOptionsLayout   = PrintOptions(yoga.PrintOptionsLayout)
	PrintOptionsStyle    = PrintOptions(yoga.PrintOptionsStyle)
	PrintOptionsChildren = PrintOptions(yoga.PrintOptionsChildren)
)

func (p PrintOptions) String() string {
	return yoga.PrintOptionsToString(yoga.PrintOptions(p))
}
