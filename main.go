package main

import (
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
)

const LOGIN = "alex_strive"

func main() {
	r := gin.Default()
	r.POST("/size2json", func(c *gin.Context) {

		file, _, err := c.Request.FormFile("image")

		defer file.Close()

		if err != nil {
			c.String(http.StatusInternalServerError, "Something went wrong")
			return
		}

		imgConfig, err := png.DecodeConfig(file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Image information decoding has failed")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"width":  imgConfig.Width,
			"height": imgConfig.Height,
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, LOGIN)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
