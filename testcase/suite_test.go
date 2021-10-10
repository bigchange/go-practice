package testcase

// Basic imports
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// suite开始测试的全局初始化可以放在此处
// 减少代码的冗余
func (s *ExampleTestSuite) SetupSuite() {
	// global initializer here
}

// suite 所有测试运行结束后才会执行
// 资源的清理和恢复
func (s *ExampleTestSuite) TearDownSuite() {
	// final cleanup operations here
}

// BeforeTest has a function to be executed right before the test
// starts and receives the suite and test names as input
// but After SetupTest, we usually don't need it
func (s *ExampleTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest, suiteName:%v,testName:%v \n", suiteName, testName)

}
// AfterTest has a function to be executed right after the test
// finishes and receives the suite and test names as input
// but after TearDownTest, we usually don't need it
func (s *ExampleTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest, suiteName:%v,testName:%v \n", suiteName, testName)
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *ExampleTestSuite) SetupTest() {
	fmt.Printf("SetupTest running \n")
	s.VariableThatShouldStartAtFive = 5
}

// run after each test
func (s *ExampleTestSuite) TearDownTest() {
	s.VariableThatShouldStartAtFive = 0
	fmt.Printf("TearDownTest running \n")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *ExampleTestSuite) TestExample() {
	s.Equal(s.VariableThatShouldStartAtFive, 5)
	assert.Equal(s.T(), 5, s.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
// 定义好测试套件入口后，在ExampleTestSuite中的TestXXX 都会被执行到
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
