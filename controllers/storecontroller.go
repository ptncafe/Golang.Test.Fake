package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStoreById(c *gin.Context) {
	idparam := c.Params.ByName("id")
	id, error := strconv.Atoi(idparam)
	if error != nil {
		// handle error
		c.JSON(400, gin.H{
			"data":     nil,
			"messages": error,
			"isError":  true,
		})
	}
	var istoreRepository IStoreRepository

	store, error := istoreRepository.GetStoreById(id)
	if error != nil {
		log.Fatal("Error creating connection pool: " + error.Error())
		c.JSON(500, gin.H{
			"data":     nil,
			"messages": error,
			"isError":  true,
		})
		return
	}

	c.JSON(200, gin.H{
		"data":     store,
		"messages": error,
		"isError":  false,
	})
	return
}
