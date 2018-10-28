package yogoa

import (
	"github.com/jackwakefield/yogoa/pkg/yoga"
)

const Undefined = float32(yoga.Undefined)

func IsUndefined(value float32) bool {
	return yoga.FloatIsUndefined(value)
}
