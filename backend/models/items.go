package models

import "gorm.io/gorm"

type Items struct {
	ID               uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name    		*string `json:"name"`
	UnitPrice       *string `json:"unit_price"`
	ItemCategory 	*string `json:"item_category"`
}


func MigrateItems(db *gorm.DB) error {
	err := db.AutoMigrate(&Items{})
	return err
}
