package tree

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type BSTTestSuite struct {
    suite.Suite
    TestTree *BinarySearchTree
}

// Make sure that TestTree is set
// before each test
func (suite *BSTTestSuite) SetupTest() {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10, 4}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    suite.TestTree = tree
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestBST(t *testing.T) {
    suite.Run(t, new(BSTTestSuite))
}

func (suite *BSTTestSuite) TestBSTInOrderWalk() {
    expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
    assert.Equal(suite.T(), expected, suite.TestTree.Root.InOrderWalk())
}

func (suite *BSTTestSuite) TestBSTPreOrderWalk() {
    expected := []interface{}{2, 1, 6, 3, 5, 4, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    assert.Equal(suite.T(), expected, suite.TestTree.Root.PreOrderWalk())
}

func (suite *BSTTestSuite) TestBSTPostOrderWalk() {
    expected := []interface{}{1, 4, 5, 3, 10, 9, 8, 11, 12, 13, 14, 15, 16, 7, 6, 2}
    assert.Equal(suite.T(), expected, suite.TestTree.Root.PostOrderWalk())
}

func (suite *BSTTestSuite) TestBSTSearch() {
    node := suite.TestTree.Root.Search(7)
    assert.Equal(suite.T(), 7, node.(*BinarySearchNode).Val)
}

func (suite *BSTTestSuite) TestBSTMinimum() {
    node := suite.TestTree.Root.Minimum()
    assert.Equal(suite.T(), 1, node.(*BinarySearchNode).Val)
}

func (suite *BSTTestSuite) TestBSTMaximum() {
    node := suite.TestTree.Root.Maximum()
    assert.Equal(suite.T(), 16, node.(*BinarySearchNode).Val)
}

func (suite *BSTTestSuite) TestSize() {
    size := suite.TestTree.Size()
    assert.Equal(suite.T(), 16, size)
}

func (suite *BSTTestSuite) TestHeight() {
}

func (suite *BSTTestSuite) TestBFS() {
}

func (suite *BSTTestSuite) TestDFS() {
}

func (suite *BSTTestSuite) TestSucessor() {
}

func (suite *BSTTestSuite) TestJoin() {
}

func (suite *BSTTestSuite) TestInvert() {
}

func (suite *BSTTestSuite) TestIsContinous() {
}

func (suite *BSTTestSuite) TestIsFoldable() {
}
