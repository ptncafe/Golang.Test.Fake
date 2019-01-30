package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"ptncafe.test/databaseprovider"
	"ptncafe.test/dtos"
)

func Get(c *gin.Context) {

	db := databaseprovider.MsSQLConnection()
	rows, error := db.Query("SELECT TOP 100 Id, Name, Code FROM Store")
	defer db.Close()
	defer rows.Close()
	if error != nil {
		log.Fatal("Error creating connection pool: " + error.Error())
		c.JSON(500, gin.H{
			"messages": error,
		})
		return
	}
	stores := []dtos.StoreDto{}

	for rows.Next() {
		var row dtos.StoreDto
		err := rows.Scan(&row.Id, &row.Name, &row.Code)
		if err != nil {
			log.Fatalf("Scan: %v", err.Error())
		}
		stores = append(stores, row)
	}

	c.JSON(200, gin.H{
		"messages": stores,
	})
	return
}
