package sort

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type SortTestSuite struct {
    suite.Suite
    TestArray, Expected []int
}

// Make sure that TestTree is set
// before each test
func (suite *SortTestSuite) SetupTest() {
    suite.TestArray = []int{15, 4, 2, 1, 3, 6, 7, 5, 12, 10, 9, 11, 14, 13, 8}
    suite.Expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
}

func TestSort(t *testing.T) {
    suite.Run(t, new(SortTestSuite))
}

func (suite *SortTestSuite) TestInsertionSort() {
    InsertionSort(suite.TestArray)
    assert.Equal(suite.T(), suite.Expected, suite.TestArray)
}

func (suite *SortTestSuite) TestMergeSort() {
    out := MergeSort(suite.TestArray)
    assert.Equal(suite.T(), suite.Expected, out)
}

func (suite *SortTestSuite) TestMergeSortInPlace() {
    MergeInPlaceSort(suite.TestArray)
    // assert.Equal(suite.T(), suite.Expected, suite.TestArray)
}

func (suite *SortTestSuite) TestQuickSort() {
    QuickSort(suite.TestArray)
    assert.Equal(suite.T(), suite.Expected, suite.TestArray)
}

func (suite *SortTestSuite) TestSelectionSort() {
    SelectionSort(suite.TestArray)
    assert.Equal(suite.T(), suite.Expected, suite.TestArray)
}

func BenchmarkInsertionSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        input := []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
        InsertionSort(input)
    }
}

func BenchmarkQuickSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        input := []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
        QuickSort(input)
    }
}

func BenchmarkMergeSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        input := []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
        MergeSort(input)
    }
}
