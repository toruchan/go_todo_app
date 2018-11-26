package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("gosample/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			//"tasks": tasks,
		})
    })
	router.Run(":8080")
    // http.HandleFunc("/", controller.IndexGET)
	// http.HandleFunc("/addTask", controller.AddTask)
    // http.ListenAndServe(":8080", nil)
	
}
