package yogoa

import "github.com/jackwakefield/yogoa/pkg/yoga"

type Edge int32

const (
	EdgeLeft       = Edge(yoga.EdgeLeft)
	EdgeTop        = Edge(yoga.EdgeTop)
	EdgeRight      = Edge(yoga.EdgeRight)
	EdgeBottom     = Edge(yoga.EdgeBottom)
	EdgeStart      = Edge(yoga.EdgeStart)
	EdgeEnd        = Edge(yoga.EdgeEnd)
	EdgeHorizontal = Edge(yoga.EdgeHorizontal)
	EdgeVertical   = Edge(yoga.EdgeVertical)
	EdgeAll        = Edge(yoga.EdgeAll)
)

func (e Edge) String() string {
	return yoga.EdgeToString(yoga.Edge(e))
}
