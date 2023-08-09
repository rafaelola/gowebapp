package models

import (
	"database/sql"
)

type Inventory struct {
	Id            int    `json:"id"`
	ProductName   string `json:"productName"`
	ProductCode   string `json:"productCode"`
	Gtin          string `json:"gtin"`
	Gs1DataMatrix string `json:"gs1DataMatrix"`
	SupplierName  string `json:"supplierName"`
	CategoryName  string `json:"categoryName"`
}

func All(db *sql.DB) ([]Inventory, error) {
	var inventories []Inventory
	results, err := db.Query("SELECT id, productName,productCode,gtin,gs1DataMatrix,supplierName,categoryName FROM inventoryMaster")
	if err != nil {
		return nil, err
	}
	defer results.Close()
	for results.Next() {
		var localInventory Inventory
		err := results.Scan(&localInventory.Id, &localInventory.ProductName,
			&localInventory.ProductCode, &localInventory.Gtin, &localInventory.Gs1DataMatrix, &localInventory.SupplierName, &localInventory.CategoryName)
		inventories = append(inventories, localInventory)
		if err != nil {
			return nil, err
		}
		if err = results.Err(); err != nil {
			return nil, err
		}

	}
	return inventories, nil
}

func One(db *sql.DB, id int) (*Inventory, error) {
	// The reason for returning a pointer to the Inventory struct instead of just the struct itself is
	//that it allows the function caller to determine if the returned data is nil (no record found) or not.
	//If the function returned the Inventory struct directly, there would be no way to distinguish between
	//a valid record and a non-existing one since both would be represented by the same zero-initialized struct.
	//By returning a pointer, the function can return nil in case of an error, indicating that no record was found,
	//while a valid record will have a non-nil pointer.
	var inventory Inventory
	err := db.QueryRow("SELECT id, productName,productCode,gtin,gs1DataMatrix,supplierName,categoryName FROM inventoryMaster WHERE id = ?", id).Scan(&inventory.Id, &inventory.ProductName,
		&inventory.ProductCode, &inventory.Gtin, &inventory.Gs1DataMatrix, &inventory.SupplierName, &inventory.CategoryName)
	if err != nil {
		return nil, err
	}
	// if no matching results found, return nil
	if inventory.Id == 0 {
		return nil, nil
	}
	return &inventory, nil
}

func Create(db *sql.DB, inventory *Inventory) error {
	result, err := db.Exec("INSERT INTO inventoryMaster(productName,productCode,gtin,gs1DataMatrix,supplierName,categoryName) VALUES(?,?,?,?,?,?)", inventory.ProductName,
		inventory.ProductCode, inventory.Gtin, inventory.Gs1DataMatrix, inventory.SupplierName, inventory.CategoryName)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	// return last inserted id as int
	inventory.Id = int(id)
	return nil
}

func Update(db *sql.DB, id int, inventory *Inventory) error {
	_, err := db.Exec("UPDATE inventoryMaster SET productName=?,productCode=?,gtin=?,gs1DataMatrix=?,supplierName=?,categoryName=? WHERE id=?", inventory.ProductName,
		inventory.ProductCode, inventory.Gtin, inventory.Gs1DataMatrix, inventory.SupplierName, inventory.CategoryName, id)
	if err != nil {
		return err
	}
	return nil

}
