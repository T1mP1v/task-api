// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE
package main

import (

    "github.com/gin-gonic/gin"

 	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "task-api/docs"
)



func main() {
	InitDB()

	r := gin.Default()

	r.Static("/frontend", "./frontend")

	r.GET("/tasks", GetTasks)
	r.GET("/tasks/:id", GetTask)
	r.POST("/tasks", CreateTask)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}
