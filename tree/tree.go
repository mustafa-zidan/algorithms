package tree

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
