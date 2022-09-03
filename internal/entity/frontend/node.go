package frontend

import "github.com/haoran-mc/gcbbs/internal/model"

type Node struct {
	model.Nodes
}

type Nodes struct {
	List []Node
}

type NodeTree struct {
	Item  Node
	Child Nodes
}
