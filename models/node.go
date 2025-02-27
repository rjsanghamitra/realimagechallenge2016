package models

type Node struct {
	Name  string
	Child map[string]*Node
}

func UpsertNode(parent *Node, name string) *Node {
	childNode, exists := parent.Child[name]
	if exists {
		return childNode
	}
	childNode = &Node{
		Name:  name,
		Child: map[string]*Node{},
	}
	parent.Child[name] = childNode
	return childNode
}
