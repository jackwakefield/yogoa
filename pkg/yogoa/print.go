package yogoa

import "github.com/jackwakefield/yogoa/pkg/yoga"

type PrintOption int32

const (
	PrintOptionLayout   = PrintOption(yoga.PrintOptionsLayout)
	PrintOptionStyle    = PrintOption(yoga.PrintOptionsStyle)
	PrintOptionChildren = PrintOption(yoga.PrintOptionsChildren)
)

func (p PrintOption) String() string {
	return yoga.PrintOptionsToString(yoga.PrintOptions(p))
}
