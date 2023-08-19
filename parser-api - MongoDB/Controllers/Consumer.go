package Controllers

import (
	"appdirs/cns-parser/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetConsumer ... Get all Consumer
func GetConsumer(c *gin.Context) {
	var consumer []Models.Consumer
	err := Models.GetAllConsumers(&consumer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, consumer)
	}
}

//CreateConsumer ... Create Consumer
func CreateConsumer(c *gin.Context) {
	var consumer Models.Consumer
	c.BindJSON(&consumer)
	err := Models.CreateConsumer(&consumer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, consumer)
	}
}

//GetConsumerByID ... Get the consumer by id
func GetConsumerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var consumer Models.Consumer
	err := Models.GetConsumerByID(&consumer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, consumer)
	}
}

//UpdateConsumer ... Update the consumer information
func UpdateConsumer(c *gin.Context) {
	var consumer Models.Consumer
	id := c.Params.ByName("id")
	err := Models.GetConsumerByID(&consumer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, consumer)
	}
	c.BindJSON(&consumer)
	err = Models.UpdateConsumer(&consumer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, consumer)
	}
}

//DeleteConsumer ... Delete the consumer
func DeleteConsumer(c *gin.Context) {
	var consumer Models.Consumer
	id := c.Params.ByName("id")
	err := Models.DeleteConsumer(&consumer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
