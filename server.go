package main

import (
	"net/http"

	"github.com/basheer-shahrour/gin-server/controller"
	"github.com/basheer-shahrour/gin-server/database"
	"github.com/basheer-shahrour/gin-server/repository"
	"github.com/basheer-shahrour/gin-server/service"
	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository(database.Database)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {

	defer database.Close()

	server := gin.Default()

	// server.Use(gindump.Dump())

	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")

	apiRoutes.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})
	apiRoutes.POST("/user", func(ctx *gin.Context) {
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Done",
			})
		}
	})
	apiRoutes.PUT("/user/:id", func(ctx *gin.Context) {
		err := userController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Done",
			})
		}
	})

	viewRoutes := server.Group("/view")

	viewRoutes.GET("/usersPage", userController.ShowAll)
	server.Run(":8000")
}
