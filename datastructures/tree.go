package main

type Node struct {
	Val      interface{}
	Parent   *Node
	Children []Node
}

func (node *Node) Delete() {
	// Remove the node from it's parents children collection
	for idx, sibling := range n.Parent.Children {
		if sibling == node {
			node.Parent.Children = append(
				node.Parent.Children[:idx],
				node.Parent.Children[idx+1:]...,
			)
		}
	}

	// If the node has any children,
	// set their parent to nil and set
	// the node's children collection to nil
	if len(node.Children) != 0 {
		for _, child := range node.Children {
			child.Parent = nil
		}
		node.Children = nil
	}
}
