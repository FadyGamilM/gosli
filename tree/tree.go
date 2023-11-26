package tree

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Tree struct {
	Root   *Node
	Lenght int
}

func (n Node) Print() string {
	return fmt.Sprintf("node : {%v}\n", n.Val)
}

func (t Tree) Print() string {
	sb := strings.Builder{}
	return sb.String()
}

// func (t Tree) InOrderTraversal(sb *strings.Builder) string {

// }

// func (t Tree) InOrderTraversalLogic(sb *strings.Builder, root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	// deal with the left part
// 	// deal with the current root node
// 	sb.WriteString(fmt.Sprintf("node : {%v}\n", root.Val))
// 	// deal with the right part

// }
