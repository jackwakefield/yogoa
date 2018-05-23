package yogoa

import "github.com/jackwakefield/yogoa/yoga"

type Wrap int32

const (
	WrapNoWrap      = Wrap(yoga.WrapNoWrap)
	WrapWrap        = Wrap(yoga.WrapWrap)
	WrapWrapReverse = Wrap(yoga.WrapWrapReverse)
)

func (w Wrap) String() string {
	return yoga.WrapToString(yoga.Wrap(w))
}
