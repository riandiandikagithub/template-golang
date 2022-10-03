package core

import (
	"io"
	"payment-simulator/internal"
	healthcheckDelivery "payment-simulator/modules/healthcheck/delivery/http"
	"payment-simulator/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	// Tracer   opentracing.Tracer
	// Reporter jaeger.Reporter
	Closer  io.Closer
	Err     error
	GinFunc gin.HandlerFunc
	Router  *gin.Engine
	// Path    map[string]string
}

var debugMode string

func init() {
	debugMode = utils.GetEnv("DEBUG_MODE", "debug")
}

// Initialize Router
func InitRouter() Router {
	return Router{}
}

func (r *Router) SetRouter() *gin.Engine {
	gin.SetMode(debugMode)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "File-Name"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	router.Use(r.GinFunc)
	router.Use(internal.GinLogger())
	router.Use(gin.Recovery())

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// docs.SwaggerInfo.Title = "OttoAG Reporting API Swagger"
	// docs.SwaggerInfo.Description = "This is a list of sample api for ottoag reporting api."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// switch utils.GetEnv("APPS_ENV", "local") {
	// case "local":
	// 	docs.SwaggerInfo.Host = "localhost:8973"
	// case "dev":
	// 	docs.SwaggerInfo.Host = "13.228.25.85:8973"
	// }
	// docs.SwaggerInfo.BasePath = header
	healthcheckDelivery.NewHealthcheckHandler(router)
	// api := router.Group("/v1")
	// api.GET("healtcheck", v1.HealtCheck)
	return router
	// router.GET(accesspointVersion, controller.VersionStatus) //new
}
