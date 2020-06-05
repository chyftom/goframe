package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Code string
	Email uint
	CreatedBy string
	UpdatedBy string
}