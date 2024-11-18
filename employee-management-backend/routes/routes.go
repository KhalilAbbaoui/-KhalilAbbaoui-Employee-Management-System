package routes

import (
    "github.com/gin-gonic/gin"
    "employee-management-system/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Define the routes
    router.POST("/login", controllers.Login)

    return router
}
