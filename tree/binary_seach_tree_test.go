package tree

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestBSTInOrderWalk(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    expected := []interface{}{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
    assert.Equal(t, expected, tree.Root.InOrderWalk())
}

func TestBSTPreOrderWalk(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    expected := []interface{}{2, 1, 6, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    assert.Equal(t, expected, tree.Root.PreOrderWalk())
}

func TestBSTPostOrderWalk(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    expected := []interface{}{1, 5, 3, 10, 9, 8, 11, 12, 13, 14, 15, 16, 7, 6, 2}
    assert.Equal(t, expected, tree.Root.PostOrderWalk())
}

func TestBSTSearch(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    node := tree.Root.Search(7)
    assert.Equal(t, 7, node.(*BinarySearchNode).Val)
}

func TestBSTMinimum(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    node := tree.Root.Minimum()
    assert.Equal(t, 1, node.(*BinarySearchNode).Val)
}

func TestBSTMaximum(t *testing.T) {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    node := tree.Root.Maximum()
    assert.Equal(t, 16, node.(*BinarySearchNode).Val)
}
