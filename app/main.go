package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.POST("/", func(ctx *gin.Context) {
		file, handler, err := ctx.Request.FormFile("binary")
		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad request")
			return
		}
		fileName := handler.Filename
		dir := os.Getenv("UPLOAD_DIR")
		out, err := os.Create(dir + "/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})

	port := os.Getenv("PORT")
	err := router.Run("0.0.0.0" + ":" + port)
	if err != nil {
		return
	}
}
