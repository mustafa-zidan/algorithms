package search

import(
	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

func (node *ds.TreeNode) BFS(item interface{}) *ds.TreeNode{

	queue := ds.NewQueue()
	queue.Push(n)
	for i := queue.Pop() ; i!= nil {
		if i.Val == item {
			return i
		}

		i.Left != nil {
			queue.Push(i.Left)
		}

		i.Right != nil {
			queue.Push(i.Right)
		}
	}
	return nil
}

func (node *ds.Node) BFS(item interface{}) *ds.Node{
	queue := ds.NewQueue()
	queue.Push(n)
	for i := queue.Pop() ; i!= nil {
		if i.Val == item {
			return i
		}
		i.Children != nil {
			queue.Push(i.Children...)
		}
	}
	return nil
}