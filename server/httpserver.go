package server

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/a-hilaly/go-ftpsd/server/handlers"
	"github.com/a-hilaly/go-ftpsd/server/middlewares"
)

//
var (
	ServerEngine *gin.Engine
	Mode         string
	DefaultPort  int
	TestPort     int
	Initilized   bool = false
)

const (
	Debug      string = "DEBUG"
	Default    string = "DEFAULT"
	Testing    string = "TESTING"
	Production string = "PRODUCTION"
)

// Init ServerEngine variable
func Init(mode string, port int, token string) {
	ServerEngine = MakeEngine()
	Mode = mode
	DefaultPort = port
	middlewares.SetToken(token)
	Initilized = true
}

// Call to run server
func Run() error {
	if Initilized == false {
		errors.New("Error Not Initialized")
	}
	return ServerEngine.Run(portString(DefaultPort))
}

// Create and assemble server items
func MakeEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	assembleHandlers(engine)
	return engine
}

func SwitchMode() {
	DefaultPort, TestPort = TestPort, DefaultPort
}

// Assemble handlers and middlewares
func assembleHandlers(g *gin.Engine) {
	//g.Use(middleware.PolicyMW())
	//g.Use(middleware.TokenizeMW())

	// Making routes
	devRoute := g.Group("/dev")
	devRoute.Use(middlewares.TokenValidationMW())
	{
		devRoute.GET("/", handlers.DevHandler)
		devRoute.GET("/test", handlers.TestHandler)
		devRoute.GET("/benchmarks", handlers.TestHandler)
		devRoute.GET("/healthcheck", handlers.HealthCheckHandler)
	}

	userRoute := g.Group("/user")
	{
		userRoute.POST("/create", handlers.CreateUserHandler)
		userRoute.POST("/update", handlers.UpdateUserHandler)
		userRoute.POST("/drop", handlers.DropUserHandler)
		userRoute.POST("/auth", handlers.AuthentificateUserHandler)
		userRoute.POST("/info", handlers.UserDataHandler)
		userRoute.POST("/stats", handlers.UserStatsHandler)
	}

	//apiConfigRoute := g.Group("/settings")
	{

	}
}

func portString(port int) string {
	return ":" + strconv.Itoa(port)
}
