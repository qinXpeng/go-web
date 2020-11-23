package gee

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	pattern  string
	part     string
	children []*TrieNode
	isWild   bool
}

func (n *TrieNode) String() string {
	return fmt.Sprintf("TrieNode{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *TrieNode) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &TrieNode{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *TrieNode) search(parts []string, height int) *TrieNode {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *TrieNode) travel(list *([]*TrieNode)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *TrieNode) matchChild(part string) *TrieNode {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *TrieNode) matchChildren(part string) []*TrieNode {
	nodes := make([]*TrieNode, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
