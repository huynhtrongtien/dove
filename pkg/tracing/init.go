package tracing

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupMiddleware(r *gin.Engine, serviceName string) {
	r.Use(otelgin.Middleware(serviceName))
}

func SetupForRouterGroup(g *gin.RouterGroup, serviceName string) {
	g.Use(otelgin.Middleware(serviceName))
}
