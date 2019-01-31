package storerepository

import (
	"fmt"
	"log"

	"ptncafe.test/databaseprovider"
	"ptncafe.test/models"
)

type IStoreRepository interface {
	GetStoreById(id int) (*models.Store, error)
}

func GetStoreById(id int) (*models.Store, error) {
	db := databaseprovider.MsSQLConnection()
	stringquery := fmt.Sprintf("SELECT  Id, Name, Code FROM Store with (Nolock) WHERE Id = %d", id)
	rows, error := db.Query(stringquery)
	defer db.Close()
	defer rows.Close()
	if error != nil {
		log.Fatal("Error creating connection pool: " + error.Error())
		return nil, error
	}
	store := models.Store{}

	for rows.Next() {
		var row models.Store
		err := rows.Scan(&row.Id, &row.Name, &row.Code)
		if err != nil {
			log.Fatalf("Scan: %v", err.Error())
		}
		store = row
	}

	return &store, nil
}
