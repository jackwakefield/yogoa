package yogoa

import "github.com/jackwakefield/yogoa/pkg/yoga"

type Wrap int32

const (
	WrapNone    = Wrap(yoga.WrapNoWrap)
	WrapWrap    = Wrap(yoga.WrapWrap)
	WrapReverse = Wrap(yoga.WrapWrapReverse)
)

func (w Wrap) String() string {
	return yoga.WrapToString(yoga.Wrap(w))
}
