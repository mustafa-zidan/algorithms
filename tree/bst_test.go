package tree

import (
    "fmt"
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
    keys := []int{8, 4, 2, 1, 3, 6, 7, 5, 12, 10, 9, 11, 14, 13, 15}
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
    expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
    assert.Equal(suite.T(), expected, suite.TestTree.Root.InOrderWalk())
}

func (suite *BSTTestSuite) TestBSTPreOrderWalk() {
    expected := []interface{}{8, 4, 2, 1, 3, 6, 5, 7, 12, 10, 9, 11, 14, 13, 15}
    assert.Equal(suite.T(), expected, suite.TestTree.Root.PreOrderWalk())
}

func (suite *BSTTestSuite) TestBSTPostOrderWalk() {
    expected := []interface{}{1, 3, 2, 5, 7, 6, 4, 9, 11, 10, 13, 15, 14, 12, 8}
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
    assert.Equal(suite.T(), 15, node.(*BinarySearchNode).Val)
}

func (suite *BSTTestSuite) TestSize() {
    size := suite.TestTree.Size()
    assert.Equal(suite.T(), 15, size)
}

func (suite *BSTTestSuite) TestHeight() {
    height := suite.TestTree.Root.Height()
    fmt.Printf("%v\n", suite.TestTree.Root.String())
    assert.Equal(suite.T(), 4, height)
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

func BenchmarkBSTInsert(b *testing.B) {
    for i := 0; i < b.N; i++ {
        keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10, 4}
        tree := &BinarySearchTree{}
        for _, key := range keys {
            tree.Insert(key)
        }
    }
}

func BenchmarkBSTSearch(b *testing.B) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10, 4}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    for i := 0; i < b.N; i++ {
        tree.Root.Search(7)
    }
}
