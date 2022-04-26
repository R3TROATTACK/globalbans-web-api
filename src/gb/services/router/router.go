package router

import "github.com/gin-gonic/gin"

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

type Router struct {
	router *gin.Engine
}

func New() *Router {
	r := gin.Default()
	return &Router{r}
}

func (r *Router) AddRoute(path string, methods []HTTPMethod, cb func(*gin.Context)) {
	for _, method := range methods {
		r.router.Handle(string(method), path, cb)
	}
}

func (r *Router) AddMiddleware(middleware func(*gin.Context)) {
	r.router.Use(middleware)
}

func (r *Router) Run(addr string) {
	r.router.Run(addr)
}

func (r *Router) Group(path string, cb gin.HandlerFunc) *gin.RouterGroup {
	return r.router.Group(path, cb)
}
