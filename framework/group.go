package framework

type IGroup interface {
	GET(string, ControllerHandler)
	POST(string, ControllerHandler)
	PUT(string, ControllerHandler)
	DELETE(string, ControllerHandler)

	Group(string) IGroup

	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core   *Core
	parent *Group
	prefix string

	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) GET(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.GET(uri, handler)
}

func (g *Group) POST(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.POST(uri, handler)
}

func (g *Group) DELETE(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.DELETE(uri, handler)
}

func (g *Group) PUT(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.PUT(uri, handler)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Group(uri string) IGroup {
	cGroup := NewGroup(g.core, uri)
	cGroup.parent = g

	return cGroup
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = middlewares
}
