package test

import (
	mockinterfce "github.com/bigchange/go-practice/testcase/test/mockgen"
	"github.com/golang/mock/gomock"
)

func (s *ExampleTestSuite) TestFoo() {
	ctrl := gomock.NewController(s.T())
	m := mockinterfce.NewMockFoo(ctrl)
	m.EXPECT().Bar(gomock.Eq(98)).DoAndReturn(func(_ int) int {
		return 104
	})
	m.Bar(98)
	m.EXPECT().Bar(gomock.Eq(99)).DoAndReturn(func(_ int) int {
		return 103
	})
	SUT(m)
}