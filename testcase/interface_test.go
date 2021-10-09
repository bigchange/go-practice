package testcase

import (
	mockinterfce "github.com/bigchange/go-practice/testcase/mockgen"
	"github.com/golang/mock/gomock"
)

func (s *ExampleTestSuite) TestFoo() {
	ctrl := gomock.NewController(s.T())
	m := mockinterfce.NewMockFoo(ctrl)
	m.EXPECT().Bar(gomock.Eq(99)).Return(103)
	SUT(m)
	m.EXPECT().Bar(gomock.Eq(99)).DoAndReturn(func(_ int) int {
		return 104
	})
	SUT(m)
}