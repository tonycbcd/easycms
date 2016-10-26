package egin

import (
    "github.com/gin-gonic/gin"
)

type Engine struct {
    *gin.Engine
}

type RouterGroup struct  {
    *gin.RouterGroup
}

type HandlerFunc func(*Context)

func NewEngine(g *gin.Engine) *Engine {
    return &Engine{Engine:g}
}

func (this *Engine) Group(relativePath string, eHandlers ...HandlerFunc) *RouterGroup {
    return &RouterGroup{RouterGroup:this.Engine.Group(relativePath, getEHandlers(eHandlers...)...)}
}

func (this *Engine) NoRoute(eHandlers ...HandlerFunc) {
    this.Engine.NoRoute(getEHandlers(eHandlers...)...)
}

func (this *Engine) PUT(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.Engine.PUT(relativePath, getEHandlers(eHandlers...)...)
}

func (this *Engine) POST(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.Engine.POST(relativePath, getEHandlers(eHandlers...)...)
}

func (this *Engine) GET(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.Engine.GET(relativePath, getEHandlers(eHandlers...)...)
}

func (this *Engine) Any(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.Engine.Any(relativePath, getEHandlers(eHandlers...)...)
}

func (this *RouterGroup) Any(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.RouterGroup.Any(relativePath, getEHandlers(eHandlers...)...)
}

func (this *RouterGroup) PUT(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.RouterGroup.PUT(relativePath, getEHandlers(eHandlers...)...)
}

func (this *RouterGroup) GET(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.RouterGroup.GET(relativePath, getEHandlers(eHandlers...)...)
}

func (this *RouterGroup) POST(relativePath string, eHandlers ...HandlerFunc) gin.IRoutes {
    return this.RouterGroup.POST(relativePath, getEHandlers(eHandlers...)...)
}

func getEHandlers(eHandlers ...HandlerFunc) []gin.HandlerFunc {
    handlers := make([]gin.HandlerFunc, len(eHandlers))
    for i := 0; i < len(eHandlers); i++ {
        eFun := eHandlers[i]
        handler := func(c *gin.Context) {
            eFun(NewContext(c))
        }
        handlers[i] = handler
    }
    return handlers
}

