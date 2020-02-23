package tree

import "fmt"

/**
 * Binray Tree Implementation with int nodes values
 * If you want to change the implementation
 *
 * the implementation needs to change
 *
 * type <TypeName> <type>
 *
 * func (t *TypeName) Less(than TypeName) bool {
 *     return // Do the comparison Here
 * }
 *
 * func (i *TypeName) Less(than TypeName) bool {
 *     return // Do the comparison Here
 * }
 *
 * type TreeNode struct {
 *     Val *TypeName
 * }
 *
 * items = []Item{1,2,3,4}
 * for _, v := range items {
 *     tree.Insert(v)
 * }
 *
 * Less and Greater should use n.Val.Less()
 *
 */

type BinarySearchTree struct {
    Root Node
}

type BinarySearchNode struct {
    Node
    Val    int
    Parent *BinarySearchNode
    Left   *BinarySearchNode
    Right  *BinarySearchNode
}

func (bst *BinarySearchTree) Insert(item interface{}) {
    if bst.Root == nil {
        bst.Root = &BinarySearchNode{Val: item.(int)}
    } else {
        bst.Root.Insert(item)
    }
}

func (bst *BinarySearchTree) Size() int {
    return bst.Root.Size()
}

func (t *BinarySearchTree) Join(other *BinarySearchTree) *BinarySearchTree {
    return nil
}

// Invert
func (t *BinarySearchTree) Invert() *BinarySearchTree {
    return nil
}

func (n *BinarySearchNode) Less(than interface{}) bool {
    return n.Val < than.(int)
}

func (n *BinarySearchNode) Greater(than interface{}) bool {
    return n.Val > than.(int)
}

func (n *BinarySearchNode) Sucessor() Node {
    return nil
}

func (n *BinarySearchNode) Insert(item interface{}) {
    if n.Less(item) {
        if n.Right != nil {
            n.Right.Insert(item)
        } else {
            n.Right = &BinarySearchNode{Val: item.(int), Parent: n}
        }
    } else if n.Greater(item) {
        if n.Left != nil {
            n.Left.Insert(item)
        } else {
            n.Left = &BinarySearchNode{Val: item.(int), Parent: n}
        }
    }
}

func (n *BinarySearchNode) InOrderWalk() []interface{} {
    output := make([]interface{}, 0)
    if n.Left != nil {
        output = append(output, n.Left.InOrderWalk()...)
    }
    output = append(output, n.Val)
    if n.Right != nil {
        output = append(output, n.Right.InOrderWalk()...)
    }
    return output
}

func (n *BinarySearchNode) PreOrderWalk() []interface{} {
    output := make([]interface{}, 0)
    output = append(output, n.Val)
    if n.Left != nil {
        output = append(output, n.Left.PreOrderWalk()...)
    }
    if n.Right != nil {
        output = append(output, n.Right.PreOrderWalk()...)
    }
    return output
}

func (n *BinarySearchNode) PostOrderWalk() []interface{} {
    output := make([]interface{}, 0)
    if n.Left != nil {
        output = append(output, n.Left.PostOrderWalk()...)
    }
    if n.Right != nil {
        output = append(output, n.Right.PostOrderWalk()...)
    }
    output = append(output, n.Val)
    return output
}

func (n *BinarySearchNode) Search(value interface{}) Node {
    if n.Val == value.(int) {
        return n
    }
    if n.Less(value) {
        return n.Right.Search(value)
    } else if n.Greater(value) {
        return n.Left.Search(value)
    }
    return nil
}

func (n *BinarySearchNode) Minimum() Node {
    if n.Left != nil {
        return n.Left.Minimum()
    }
    return n
}

func (n *BinarySearchNode) Maximum() Node {
    if n.Right != nil {
        return n.Right.Maximum()
    }
    return n
}

func (n *BinarySearchNode) String() string {
    return fmt.Sprintf("{Val: %d, Left: %v, Right:%v}", n.Val, n.Left, n.Right)
}

func (n *BinarySearchNode) Size() int {
    if n == nil {
        return 0
    }
    return 1 + n.Left.Size() + n.Right.Size()
}

func (n *BinarySearchNode) Height() int {
    return 0
}

func (n *BinarySearchNode) BFS(value interface{}) Node {
    return n
}

func (n *BinarySearchNode) DFS(item interface{}) Node {
    return n
}

// A tree is Continuous tree if in each root to leaf path,
// absolute difference between keys of two adjacent is 1.
// We are given a binary tree, we need to check if tree is continuous or not.
func (n *BinarySearchNode) IsContinous() bool {
    return false
}

// A tree can be folded if left and right subtrees of the tree are structure
// wise mirror image of each other. An empty tree is considered as foldable.
func (n *BinarySearchNode) IsFoldable() bool {
    return false
}
