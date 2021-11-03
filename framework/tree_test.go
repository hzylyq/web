package framework

import "testing"

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: nil,
		children: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: nil,
				children: nil,
			},
			{
				isLast:   false,
				segment:  ":id",
				handlers: nil,
				children: nil,
			},
		},
	}

	{
		nodes := root.filterChildNode("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChildNode(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}

}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: nil,
		children: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: nil,
				children: []*node{
					{
						isLast:   true,
						segment:  "BAR",
						handlers: nil,
						children: []*node{},
					},
				},
			},
			{
				isLast:   true,
				segment:  ":id",
				handlers: nil,
				children: nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}
}
