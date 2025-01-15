package entity

import (
	"gorm.io/gorm"
)

type Products struct{
	gorm.Model 
	Name string `valid:"required~Name is required"`
	Price float64
	SKU string
}
