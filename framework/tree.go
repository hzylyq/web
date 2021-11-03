package framework

import (
	"fmt"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast   bool
	segment  string
	children []*node

	handlers []ControllerHandler
}

func newNode() *node {
	return &node{
		isLast:   false,
		segment:  "",
		children: make([]*node, 0),
	}
}

func NewTree() *Tree {
	return &Tree{
		root: newNode(),
	}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNode(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.children
	}

	nodes := make([]*node, 0, len(n.children))

	for _, child := range n.children {
		if isWildSegment(child.segment) {
			nodes = append(nodes, child)
		} else if child.segment == segment {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	children := n.filterChildNode(segment)
	if len(children) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, child := range children {
			if child.isLast {
				return child
			}
		}

		return nil
	}

	for _, child := range children {
		matchNode := child.matchNode(segments[1])
		if matchNode != nil {
			return matchNode
		}
	}

	return nil
}

func (tree *Tree) AddRouter(uri string, handler []ControllerHandler) error {
	n := tree.root

	if n.matchNode(uri) != nil {
		return fmt.Errorf("root exist:%s", uri)
	}

	segments := strings.Split(uri, "/")

	for idx, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		isLast := idx == len(segments)-1

		var objNode *node

		children := n.filterChildNode(segment)
		if len(children) > 0 {
			for _, child := range children {
				if child.segment == segment {
					objNode = child
					break
				}
			}
		}

		if objNode == nil {
			child := newNode()
			child.segment = segment
			if isLast {
				child.isLast = true
				child.handlers = handler
			}
			n.children = append(n.children, child)
			objNode = child
		}

		n = objNode
	}

	return nil
}

func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}
