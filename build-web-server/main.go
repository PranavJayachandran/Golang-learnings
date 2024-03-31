package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}
func getDataByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range data {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not dound"})
}
func postData(c *gin.Context) {
	var newData Person
	if err := c.BindJSON(&newData); err != nil {
		fmt.Println(err)
		return
	}

	data = append(data, newData)
	c.IndentedJSON(http.StatusCreated, newData)
}
func main() {
	router := gin.Default()
	router.GET("/data", getData)
	router.GET("/data/:id", getDataByID)
	router.POST("/data", postData)
	router.Run("localhost:8000")
}
