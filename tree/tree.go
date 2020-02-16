package tree

type Tree interface {
    InOrder() []interface{}
    PreOrder() []interface{}
    PostOrder() []interface{}
    Search() *Node
    Insert(item interface{})
    Minimum() *Node
    Maximum() *Node
    Replace(src, dest *Node)
    Delete(item *Node)
    Size() int
    BFS(item interface{}) *Node
    DFS(item interface{}) *Node
}

type Node interface {
    Sucessor() *Node
    InOrderWalk() []interface{}
    PreOrderWalk() []interface{}
    PostOrderWalk() []interface{}
    Insert(item interface{})
    Less(than interface{}) bool
    Greater(than interface{}) bool
}
