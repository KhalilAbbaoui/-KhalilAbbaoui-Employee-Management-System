package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "employee-management-system/config"
    "employee-management-system/controllers"
    "employee-management-system/middlewares"
    "github.com/gin-contrib/cors"
)

func main() {
    // Initialize database connection
    if err := config.InitDB(); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    // Create a new Gin router
    r := gin.Default()

    // Enable CORS middleware for the frontend (Angular on localhost:4200)
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8081"}, // Allow Angular front-end to access
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Routes API
    api := r.Group("/api")
    api.Use(middlewares.JWTAuthMiddleware()) // Protect these routes with JWT authentication
    {
        api.GET("/employees", controllers.GetEmployees)
        api.GET("/employees/:id", controllers.GetEmployee)
        api.POST("/employees", controllers.CreateEmployee)
        api.PUT("/employees/:id", controllers.UpdateEmployee)
        api.DELETE("/employees/:id", controllers.DeleteEmployee)
    }

    // Login route (no JWT middleware needed)
    r.POST("/api/employees/login", controllers.Login)

    // Start the server on port 8080
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
