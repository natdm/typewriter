package template

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidEcmaTestSuite struct {
	suite.Suite
}

func TestValidEcmaTestSuite(t *testing.T) {
	suite.Run(t, new(ValidEcmaTestSuite))
}

func (s *ValidEcmaTestSuite) TestValidEcma() {
	s.True(propertyShouldBeQuoted("hello-world"))
	s.True(propertyShouldBeQuoted("hello#world"))
	s.True(propertyShouldBeQuoted("你好世界"))
	s.True(propertyShouldBeQuoted("hello/world"))
	s.False(propertyShouldBeQuoted("$helloWorld"))
	s.False(propertyShouldBeQuoted("helloWorld"))
	s.False(propertyShouldBeQuoted("hello_world"))
}
