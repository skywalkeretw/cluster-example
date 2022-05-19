package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"body": "Hello From Go",
	})
}

func callPython(c *gin.Context) {
	r, err := http.Get("http://python-server:8080")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	var data interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, data)
}

func callNode(c *gin.Context) {
	r, err := http.Get("http://node-server:8080")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	var data interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/node", callNode)
	r.GET("/python", callPython)
	r.Run(":8080")
}
