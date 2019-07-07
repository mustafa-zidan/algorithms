package search

import(
	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

func (node *ds.TreeNode) DFS(item interface{}) *ds.TreeNode{
	stack := ds.NewStack()
	stack.Push(n)
	for i := stack.Pop() ; i!= nil {
		if i.Val == item {
			return i
		}

		i.Left != nil {
			stack.Push(i.Left)
		}

		i.Right != nil {
			stack.Push(i.Right)
		}
	}
	return nil
}

func (node *ds.Node) DFS(item interface{}) *ds.Node{
	stack := ds.NewStack()
	stack.Push(n)
	for i := stack.Pop() ; i!= nil {
		if i.Val == item {
			return i
		}
		i.Children != nil {
			stack.Push(i.Children...)
		}
	}
	return nil
}