package server


import (
    "errors"
    "strconv"

    "github.com/gin-gonic/gin"

    "github.com/a-hilaly/supfile-api/server/handlers"
    //"github.com/a-hilaly/supfile-api/server/middleware"
)

//
var (
    ServerEngine *gin.Engine
    DefaultPort          int
    TestPort             int
    Initilized          bool = false
)

// Init ServerEngine variable
func Init(port int) {
    ServerEngine = MakeEngine()
    DefaultPort = port
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
    {
        devRoute.GET("/", handlers.DevHandler)
        devRoute.GET("/test", handlers.TestHandler)
        devRoute.GET("/benchmarks", handlers.TestHandler)
        devRoute.GET("/healthcheck", handlers.HealthCheckHandler)
    }

    userRoute := g.Group("/user")
    {
        userRoute.GET("/create", handlers.CreateUserHandler)
        userRoute.GET("/update", handlers.UpdateUserHandler)
        userRoute.GET("/drop", handlers.DropUserHandler)
        userRoute.GET("/auth", handlers.AuthentificateUserHandler)
        userRoute.GET("/info", handlers.UserDataHandler)
        userRoute.GET("/stats", handlers.UserStatsHandler)
    }

    apiConfigRoute := g.Group("/settings")
    {

    }
}

func portString(port int) string {
    return ":" + strconv.Itoa(port)
}
