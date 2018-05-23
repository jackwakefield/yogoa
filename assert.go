package yogoa

import (
	"github.com/jackwakefield/yogoa/yoga"
)

func Assert(condition bool, message string) {
	yoga.Assert(condition, message)
}
