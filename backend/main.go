package main

import (
	"github.com/tonphaii/sa-65-example/controller"

	"github.com/tonphaii/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes

	r.GET("/car", controller.ListCars)
	r.GET("/car/:id", controller.GetCar)
	r.POST("/car", controller.CreateCar)
	r.PATCH("/car", controller.UpdateCar)
	r.DELETE("/car/:id", controller.DeleteRecordTimeOut)

	r.GET("/case", controller.ListCases)
	r.GET("/case/:id", controller.GetCase)
	r.POST("/case", controller.CreateCase)
	r.PATCH("/case", controller.UpdateCase)
	r.DELETE("/case/:id", controller.DeleteCase)

	r.GET("/status", controller.ListStatus)
	r.GET("/status/:id", controller.GetStatus)
	r.POST("/status", controller.CreateStatus)
	r.PATCH("/status", controller.UpdateStatus)
	r.DELETE("/status/:id", controller.DeleteStatus)

	r.GET("/recordtimeout", controller.ListRecordTimeOuts)
	r.GET("/recordtimeout/:id", controller.GetRecordTimeOut)
	r.POST("/recordtimeout", controller.CreateRecordTimeOut)
	r.PATCH("/recordtimeout", controller.UpdateRecordTimeOut)
	r.DELETE("/recordtimeout/:id", controller.DeleteRecordTimeOut)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
