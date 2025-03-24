package routers

import (
	"voting/internal/application/usecase"
	"voting/internal/infraestructure/primary/handlers/handlers"
	"voting/internal/infraestructure/primary/handlers/middleware"
	"voting/internal/infraestructure/secundary/postgres/connectpostgres"
	"voting/internal/infraestructure/secundary/postgres/repository"
	"voting/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	log := logger.NewLogger()
	router := gin.New()
	router.Use(gin.RecoveryWithWriter(log.Writer()))
	router.Use(middleware.GinLogger(log))

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dbConnection := connectpostgres.New()
	repo := repository.New(dbConnection)
	usecase := usecase.New(repo)
	hanlder := handlers.New(usecase)

	api := router.Group("/api")
	{
		api.POST("/vote-group", hanlder.CreateVoteGroup)
		api.POST("/create-user", hanlder.CreateUser)
		api.POST("/create-house", hanlder.CreateHouse)
		api.GET("/get-users/:groupID", hanlder.GetUsers)
		api.GET("/get-houses", hanlder.GetHouses)
		api.GET("/get-groups", hanlder.GetGroups)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
