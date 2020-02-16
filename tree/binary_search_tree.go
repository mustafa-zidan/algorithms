package tree

import "fmt"

/**
 * Binray Tree Implementation with int nodes values
 * If you want to change the implementation
 *
 * the implementation needs to change
 *
 * type Int int
 *
 * func (i *Int) Less(than Int) bool {
 *     return i < than
 * }
 *
 * func (i *Int) Less(than Int) bool {
 *     return i < than
 * }
 *
 * type TreeNode struct {
 *     Val *Int
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

func (bst *BinarySearchTree) Minimum() Node {
    return bst.Root.Minimum()
}

func (bst *BinarySearchTree) Maximum() Node {
    return bst.Root.Maximum()
}

func (bst *BinarySearchTree) Replace(src, dest *BinarySearchNode) {

}

func (bst *BinarySearchTree) Delete(item *BinarySearchNode) {

}

func (bst *BinarySearchTree) Size() int {
    return 0
}

func (bst *BinarySearchTree) BFS(val interface{}) Node {
    return nil
}

func (bst *BinarySearchTree) DFS(val interface{}) Node {
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
    return fmt.Sprintf("%d\n", n.Val)
}
