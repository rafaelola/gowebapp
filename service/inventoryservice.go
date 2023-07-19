package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

var db *sql.DB

type Inventory struct {
	Id            int    `json:"id"`
	ProductName   string `json:"productName"`
	ProductCode   string `json:"productCode"`
	Gtin          string `json:"gtin"`
	Gs1DataMatrix string `json:"gs1DataMatrix"`
	SupplierName  string `json:"supplierName"`
	CategoryName  string `json:"categoryName"`
}

func GetInventory() []byte {
	// connect to db
	// query the db
	var inventories []Inventory
	results, err := db.Query("SELECT id, productName,productCode,gtin,gs1DataMatrix,supplierName,categoryName FROM inventoryMaster")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for results.Next() {
		var localInventory Inventory
		err = results.Scan(&localInventory.Id, &localInventory.ProductName,
			&localInventory.ProductCode, &localInventory.Gtin, &localInventory.Gs1DataMatrix, &localInventory.SupplierName, &localInventory.CategoryName)
		inventories = append(inventories, localInventory)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		err := results.Close()
		if err != nil {
			return nil
		}

	}
	jsonInventories, err := json.Marshal(inventories)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonInventories
}
