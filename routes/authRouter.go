package routes

import (
	"github.com/DEVunderdog/monolithic-notes-app/controllers"
	"github.com/DEVunderdog/monolithic-notes-app/middleware"
	"github.com/gin-gonic/gin"
)


func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("users/notes/create", middleware.Authenticate(),controllers.CreateNote())
}