package datastructures

type BinarySearchTree struct {
	Root *TreeNode
}

type TreeNode struct {
	Val    int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

func (t *BinarySearchTree) InOrderTreeWalk() []int {
	return t.Root.InOrderTreeWalk()
}

func (t *BinarySearchTree) PreOrderTreeWalk() []int {
	return t.Root.PreOrderTreeWalk()
}

func (t *BinarySearchTree) PostOrderTreeWalk() []int {
	return t.Root.PostOrderTreeWalk()
}

func (n *TreeNode) InOrderTreeWalk() []int {
	result := make([]int, 0)
	if n.Left != nil {
		result = append(result, n.Left.InOrderTreeWalk()...)
	}
	result = append(result, n.Val)
	if n.Right != nil {
		result = append(result, n.Right.InOrderTreeWalk()...)
	}
	return result
}

func (n *TreeNode) PreOrderTreeWalk() []int {
	result := make([]int, 0)
	result = append(result, n.Val)
	if n.Left != nil {
		result = append(result, n.Left.PreOrderTreeWalk()...)
	}
	if n.Right != nil {
		result = append(result, n.Right.PreOrderTreeWalk()...)
	}
	return result
}

func (n *TreeNode) PostOrderTreeWalk() []int {
	result := make([]int, 0)
	if n.Left != nil {
		result = append(result, n.Left.PostOrderTreeWalk()...)
	}

	if n.Right != nil {
		result = append(result, n.Right.PostOrderTreeWalk()...)
	}
	result = append(result, n.Val)
	return result
}

func (t *BinarySearchTree) Minimum() *TreeNode {
	min := t.Root
	for min.Left != nil {
		min = min.Left
	}
	return min
}

func (n *TreeNode) Minimum() *TreeNode {
	min := n
	for min.Left != nil {
		min = min.Left
	}
	return min
}

func (n *TreeNode) Sucessor() *TreeNode {
	if n.Right != nil {
		return n.Minimum()
	}
	sucessor := n.Parent
	for sucessor != nil && *n == *sucessor.Right {
		n = sucessor
		sucessor = sucessor.Parent
	}
	return sucessor
}

func (t *BinarySearchTree) Insert(item int) {
	if t.Root == nil {
		t.Root = &TreeNode{item, nil, nil, nil}
	} else {
		t.Root.Insert(item)
	}
}

func (n *TreeNode) Insert(item int) {
	if item > n.Val {
		if n.Right != nil {
			n.Right.Insert(item)
		} else {
			n.Right = &TreeNode{item, n, nil, nil}
		}
	} else {
		if n.Left != nil {
			n.Left.Insert(item)
		} else {
			n.Left = &TreeNode{item, n, nil, nil}
		}
	}
}

func (t *BinarySearchTree) Search(item int) *TreeNode {
	return t.Root.Search(item)
}

func (n *TreeNode) Search(item int) *TreeNode {
	if item == n.Val {
		return n
	} else if item > n.Val && n.Right != nil {
		return n.Right.Search(item)
	} else if n.Left != nil {
		return n.Left.Search(item)
	}
	return nil
}

func (t *BinarySearchTree) Replace(src, dest *TreeNode) {
	if src.Parent == nil {
		t.Root = dest
	} else if src.Parent.Left == src {
		src.Parent.Left = dest
	} else {
		src.Parent.Right = dest
	}
	dest.Parent = src.Parent
}

func (t *BinarySearchTree) Delete(node *TreeNode) {
	if node.Left == nil {
		t.Replace(node, node.Right)
	} else if node.Right == nil {
		t.Replace(node, node.Left)
	} else {
		successor := node.Right.Minimum()
		if successor.Parent != node {
			t.Replace(successor, node.Right)
		}
		t.Replace(node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
	}
}

func (t *BinarySearchTree) Size() int {
	root := t.Root
	return root.Size()
}

func (n *TreeNode) Size() int {
	if n.Left == nil && n.Right == nil {
		return 1
	} else if n.Left == nil {
		return 1 + n.Right.Size()
	} else if n.Right == nil {
		return 1 + n.Left.Size()
	}
	return 1 + n.Left.Size() + n.Right.Size()
}

// A tree is Continuous tree if in each root to leaf path,
// absolute difference between keys of two adjacent is 1.
// We are given a binary tree, we need to check if tree is continuous or not.
func (n *TreeNode) IsContinous() bool {
	if n == nil {
		return true
	}
	valid := true
	if n.Left != nil {
		valid = valid && (n.Val-n.Left.Val) == 1
	}
	if n.Right != nil {
		valid = valid && (n.Right.Val-n.Val) == 1
	}
	return valid && n.Left.IsContinous() && n.Right.IsContinous()
}

// A tree can be folded if left and right subtrees of the tree are structure
// wise mirror image of each other. An empty tree is considered as foldable.
func (n *TreeNode) IsFoldable() bool {
	return false
}

func (t *BinarySearchTree) Join(other *BinarySearchTree) *BinarySearchTree {
	return nil
}

// invert
func (t *BinarySearchTree) Invert() *BinarySearchTree {
	return nil
}
