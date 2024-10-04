package server

import (
	"database/sql"
	"log"
	"bus_booking/controllers"
	"bus_booking/repositories"
	"bus_booking/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	AuthController  *controllers.AuthController
	BusesController *controllers.BusController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
    // Initialize repositories and services
    userRepository := repositories.NewUserRepository(dbHandler)
    busRepository := repositories.NewBusRepository(dbHandler)
    orderRepository := repositories.NewOrderRepository(dbHandler)
    ratingRepository := repositories.NewRatingRepository(dbHandler)

    authService := services.NewAuthService(userRepository)
    busService := services.NewBusService(busRepository)
    orderService := services.NewOrderService(orderRepository, busRepository)
    ratingService := services.NewRatingService(ratingRepository)

    authController := controllers.NewAuthController(authService)
    busController := controllers.NewBusController(busService)

    // Setup router
    router := gin.Default()

    // // Routes for user login
    router.POST("/login", authController.Login)

    // // Routes for buses
    router.POST("/buses", busController.CreateBus)
    router.GET("/buses", busController.GetAllBuses)
    router.GET("/buses/:id", busController.GetBusByID)
    router.PUT("/buses/:id", busController.UpdateBus)
    router.DELETE("/buses/:id", busController.DeleteBus)

	return HttpServer{
		config:            config,
		router:            router,
		authController: authController,
		busController: busController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}


// package server

// import (
// 	"database/sql"
// 	"log"
// 	"bus_booking/controllers"
// 	"bus_booking/repositories"
// 	"bus_booking/services"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// )

// type HttpServer struct {
// 	config            *viper.Viper
// 	router            *gin.Engine
// 	BusesController *controllers.BusesController
// }

// func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
//     // Initialize repositories and services
//     userRepository := repositories.NewUserRepository(db)
//     busRepository := repositories.NewBusRepository(db)
//     orderRepository := repositories.NewOrderRepository(db)
//     ratingRepository := repositories.NewRatingRepository(db)

//     authService := services.NewAuthService(userRepository)
//     busService := services.NewBusService(busRepository)
//     orderService := services.NewOrderService(orderRepository, busRepository)
//     ratingService := services.NewRatingService(ratingRepository)

//     authController := controllers.NewAuthController(authService)
//     busController := controllers.NewBusController(busService)

//     // Setup router
//     router := gin.Default()

//     // Routes for user login
//     router.POST("/login", authController.Login)

//     // Routes for buses
//     router.POST("/buses", busController.CreateBus)
//     router.GET("/buses", busController.GetAllBuses)
//     router.GET("/buses/:id", busController.GetBusByID)
//     router.PUT("/buses/:id", busController.UpdateBus)
//     router.DELETE("/buses/:id", busController.DeleteBus)

//     // Start server
//     if err := router.Run(viper.GetString("http.server_address")); err != nil {
//         log.Fatalf("Error starting server: %v", err)
//     }
// }



















// package server

// import (
//     "database/sql"
//     "log"
//     "bus_booking/controllers"
//     "bus_booking/repositories"
//     "bus_booking/services"
//     "github.com/gin-gonic/gin"
//     "github.com/spf13/viper"
//     _ "github.com/go-sql-driver/mysql"
// )

// type HttpServer struct {
//     config              *viper.Viper
//     router              *gin.Engine
//     authController      *controllers.AuthController
//     busController       *controllers.BusController
// }

// func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
//     userRepository := repositories.NewUserRepository(dbHandler)
//     busRepository := repositories.NewBusRepository(dbHandler)
//     orderRepository := repositories.NewOrderRepository(dbHandler)
//     ratingRepository := repositories.NewRatingRepository(dbHandler)

//     authService := services.NewAuthService(userRepository)
//     busService := services.NewBusService(busRepository)

//     authController := controllers.NewAuthController(authService)
//     busController := controllers.NewBusController(busService)

//     router := gin.Default()

//     // Define your routes here
//     router.POST("/login", authController.Login)
//     router.POST("/buses", busController.CreateBus)
//     router.GET("/buses", busController.GetAllBuses)
//     router.GET("/buses/:id", busController.GetBusByID)
//     router.PUT("/buses/:id", busController.UpdateBus)
//     router.DELETE("/buses/:id", busController.DeleteBus)

//     return HttpServer{
//         config:         config,
//         router:         router,
//         authController: authController,
//         busController:  busController,
//     }
// }

// func (hs HttpServer) Start() {
//     err := hs.router.Run(hs.config.GetString("http.server_address"))
//     if err != nil {
//         log.Fatalf("Error while starting HTTP server: %v", err)
//     }
// }











// package server

// import (
//     "database/sql"
//     "log"
//     "bus_booking/controllers"
//     "bus_booking/repositories"
//     "bus_booking/services"
//     "github.com/gin-gonic/gin"
//     "github.com/spf13/viper"
//     _ "github.com/go-sql-driver/mysql"
// )

// type HttpServer struct {
//     config         *viper.Viper
//     router         *gin.Engine
//     authController *controllers.AuthController
//     busController  *controllers.BusController
// }

// func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
//     userRepository := repositories.NewUserRepository(dbHandler)
//     busRepository := repositories.NewBusRepository(dbHandler)

//     authService := services.NewAuthService(userRepository)
//     busService := services.NewBusService(busRepository)

//     authController := controllers.NewAuthController(authService)
//     busController := controllers.NewBusController(busService)

//     router := gin.Default()

//     // Define your routes here
//     router.POST("/login", authController.Login)
//     router.POST("/buses", busController.CreateBus)
//     router.GET("/buses", busController.GetAllBuses)
//     router.GET("/buses/:id", busController.GetBusByID)
//     router.PUT("/buses/:id", busController.UpdateBus)
//     router.DELETE("/buses/:id", busController.DeleteBus)

//     return HttpServer{
//         config:         config,
//         router:         router,
//         authController: authController,
//         busController:  busController,
//     }
// }

// func (hs HttpServer) Start() {
//     err := hs.router.Run(hs.config.GetString("http.server_address"))
//     if err != nil {
//         log.Fatalf("Error while starting HTTP server: %v", err)
//     }
// }
