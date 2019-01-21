package main

import (
	"../controller"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("src/view/*")

	router.GET("/", func(c *gin.Context) {
		todos := controller.GetAll()

		//todos := ctrl.GetAll()

		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"todos": todos,
		})
    })

	router.POST("/", func(c *gin.Context) {
		text := c.PostForm("task")
		controller.AddTask(text)
		todos := controller.GetAll()
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"todos": todos,
		})
	})

	router.POST("/hoge", func(c *gin.Context) {
		//text := c.PostForm("task")
		controller.AddTask("deldel")
		todos := controller.GetAll()
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"todos": todos,
		})
	})

	
	router.Run(":8080")
}
