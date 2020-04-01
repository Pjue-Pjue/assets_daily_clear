package main

import (
	"Pjue-Pjue/assets_daily_clear_web/backend/controllers"
	"github.com/gin-gonic/gin"
	"log"
)


func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(CORS())

	//r.POST("/api/login", controllers.Login)
	r.POST("/api/addAccount", controllers.AddAccount)
	r.GET("/api/getAccounts", controllers.GetAccounts)
	r.POST("/api/updateAccount", controllers.UpdateAccount)


	log.Println("Server is ready.")
	r.Run(":3001")
}

// CORS CORS handle function
func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}
