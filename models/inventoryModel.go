package models

type InventoryModel struct {
	Id            int    `json:"id"`
	ProductName   string `json:"productName"`
	ProductCode   string `json:"productCode"`
	Gtin          string `json:"gtin"`
	Gs1DataMatrix string `json:"gs1DataMatrix"`
	SupplierName  string `json:"supplierName"`
	CategoryName  string `json:"categoryName"`
}
