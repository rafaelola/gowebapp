package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

type InventoryModel struct {
	Id            int    `json:"id"`
	ProductName   string `json:"productName" validate:"required,min=3,max=12"`
	ProductCode   string `json:"productCode"`
	Gtin          string `json:"gtin" validate:"required,min=3,max=12"`
	Gs1DataMatrix string `json:"gs1DataMatrix" validate:"required,min=3,max=12"`
	SupplierName  string `json:"supplierName"`
	CategoryName  string `json:"categoryName"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

func (i *InventoryModel) Validate() error {
	validate = validator.New()
	errs := validate.Struct(i)
	if errs != nil {
		var errors []*IError
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, &IError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
		result, _ := json.Marshal(errors)
		resultString := string(result)
		return fmt.Errorf(resultString)
	}
	return nil
}

/*func (i *InventoryModel) Validate() error {
	validate = validator.New()
	errs := validate.Struct(i)
	if errs != nil {
		var validationErrs []string
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrs = append(validationErrs, fmt.Sprintf("Field: %s, Tag: %s, Value: %s", err.Field(), err.Tag(), err.Value()))

		}
		return fmt.Errorf(strings.Join(validationErrs, ", "))
	}
	return nil
}*/
