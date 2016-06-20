package echoic

type (
	Group struct {
		echoic Echoic
	}
)

func (g *Group) Use(m ...Middleware) {
	for _, h := range m {
		g.echoic.middleware = append(g.echoic.middleware, wrapMiddleware(h))
	}
}

func (g *Group) Connect(path string, h Handler) {
	g.echoic.Connect(path, h)
}

func (g *Group) Delete(path string, h Handler) {
	g.echoic.Delete(path, h)
}

func (g *Group) Get(path string, h Handler) {
	g.echoic.Get(path, h)
}

func (g *Group) Head(path string, h Handler) {
	g.echoic.Head(path, h)
}

func (g *Group) Options(path string, h Handler) {
	g.echoic.Options(path, h)
}

func (g *Group) Patch(path string, h Handler) {
	g.echoic.Patch(path, h)
}

func (g *Group) Post(path string, h Handler) {
	g.echoic.Post(path, h)
}

func (g *Group) Put(path string, h Handler) {
	g.echoic.Put(path, h)
}

func (g *Group) Trace(path string, h Handler) {
	g.echoic.Trace(path, h)
}

func (g *Group) Static(path, root string) {
	g.echoic.Static(path, root)
}

func (g *Group) ServeDir(path, root string) {
	g.echoic.ServeDir(path, root)
}

func (g *Group) ServeFile(path, file string) {
	g.echoic.ServeFile(path, file)
}

func (g *Group) Group(prefix string, m ...Middleware) *Group {
	return g.echoic.Group(prefix, m...)
}
