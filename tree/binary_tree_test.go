package tree

import (
    "testing"

    "github.com/stretchr/testify/suite"
)

type BTTestSuite struct {
    suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestBT(t *testing.T) {
    suite.Run(t, new(BTTestSuite))
}

func (suite *BSTTestSuite) TestMaxPathSum() {
    // tests := []struct {
    //     input    []interface{}
    //     expected int
    // }{
    //     {[]interface{}{1, 2, 3}, 6},
    //     {[]interface{}{-10, 9, 20, nil, nil, 15, 7}, 42},
    // }
    // for _, t := range tests {
    //     tree := &BinaryTree{}
    //     tree.Build(t.input)
    //     assert.Equal(suite.T(), t.expected, tree.Root.MaxPathSum())
    // }
}
