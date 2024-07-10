package prometheus

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Config struct {
	Env string
	Namespace string
	ServiceName string
}

func Setup(r *gin.Engine, config *Config) *ginprom.Prometheus {
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Path("/metrics"),
		ginprom.Namespace(config.Env+"_"+config.Namespace),
		ginprom.Subsystem(config.ServiceName),
	)
	r.Use(p.Instrument())

	return p
}

func SetupForRouterGroup(r *gin.Engine, g *gin.RouterGroup, config *Config) *ginprom.Prometheus {
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Path("/metrics"),
		ginprom.Namespace(config.Env+"_"+config.Namespace),
		ginprom.Subsystem(config.ServiceName),
	)

	g.Use(otelgin.Middleware(config.ServiceName))
	g.Use(p.Instrument())

	return p
}
