package yogoa

import (
	"github.com/jackwakefield/yogoa/pkg/yoga"
)

type Experiment int32

const (
	ExperimentWebFlexBasis = Experiment(yoga.ExperimentalFeatureWebFlexBasis)
)

func (e Experiment) String() string {
	return yoga.ExperimentalFeatureToString(yoga.ExperimentalFeature(e))
}
