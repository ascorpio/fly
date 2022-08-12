package framework

// IGroup 代表前缀分组
type IGroup interface {
	Group(string) IGroup
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

// Group 实现 IGroup 接口
type Group struct {
	core   *Core
	prefix string
}

// NewGroup 初始化
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (g *Group) Group(prefix string) IGroup {
	return NewGroup(g.core, g.prefix+prefix)
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}
