package heap

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type HeapTestSuite struct {
    suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestHeap(t *testing.T) {
    suite.Run(t, new(HeapTestSuite))
}
func (suite *HeapTestSuite) TestNewMaxHeap() {
    h := NewMaxHeap(1, 2, 3, 4, 5)
    assert.Equal(suite.T(), 5, h.Size())
    assert.Equal(suite.T(), 5, (*h)[0])
}
