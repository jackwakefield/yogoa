package yogoa

import (
	"github.com/jackwakefield/yogoa/pkg/yoga"
)

type Display int32

const (
	DisplayFlex = Display(yoga.DisplayFlex)
	DisplayNone = Display(yoga.DisplayNone)
)

func (d Display) String() string {
	return yoga.DisplayToString(yoga.Display(d))
}
