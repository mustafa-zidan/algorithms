package tree

type Tree interface {
    InOrder() []interface{}
    PreOrder() []interface{}
    PostOrder() []interface{}
    Search(item interface{}) Node
    Insert(item interface{})
    Minimum() Node
    Maximum() Node
    Replace(src, dest *Node)
    Delete(item *Node)
    Size() int
    BFS(item interface{}) Node
    DFS(item interface{}) Node
}

type Node interface {
    Sucessor() Node
    Search(item interface{}) Node
    InOrderWalk() []interface{}
    PreOrderWalk() []interface{}
    PostOrderWalk() []interface{}
    Insert(item interface{})
    Minimum() Node
    Maximum() Node
    Less(than interface{}) bool
    Greater(than interface{}) bool
    String() string
}
