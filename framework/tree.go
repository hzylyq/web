package framework

import "strings"

type Tree struct {
	root *node
}

type node struct {
	isLast   bool
	segment  string
	handler  ControllerHandler
	children []*node
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

func isWildSegment(segment string ) bool {
	return strings.HasPrefix(segment, ":")
}

func
