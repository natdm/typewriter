package template

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FuncmapTestSuite struct {
	suite.Suite
}

func TestFuncmapTestSuite(t *testing.T) {
	suite.Run(t, new(FuncmapTestSuite))
}

func (s *FuncmapTestSuite) TestFlowTypes() {
	for i := 0; i < len(conversions[Flow]); i = i + 2 {
		s.Equal(conversions[Flow][i+1], updateTypes(conversions[Flow])(conversions[Flow][i]))
	}
}
