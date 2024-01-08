package models

import "gorm.io/gorm"

type Invoices struct {
	ID          uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	MobileNo    *string `json:"mobile_no"`
	Email       *string `json:"email"`
	Address     *string `json:"address"`
	BillingType *string `json:"billing_type"`
}


func MigrateInvoices(db *gorm.DB) error {
	err := db.AutoMigrate(&Invoices{})
	return err
}
