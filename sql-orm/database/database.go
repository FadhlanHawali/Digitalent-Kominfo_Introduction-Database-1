package database

import (
	"gorm.io/gorm"
	"log"
)

type CustomerORM struct {
	ID int `gorm:"primary_key" json:"-"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	NpwpId string `json:"npwp_id"`
	Age int `json:"age"`
	CustomerType string `json:"customer_type"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	PhoneNumber string `json:"phone_number"`
	AccountORM []AccountORM `gorm:"ForeignKey:IdCustomerRefer";json:"account_orm"`
}

type AccountORM struct {
	ID int `gorm:"primary_key" json:"-"`
	IdCustomerRefer int `json:"-"`
	Balance int `json:"balance"`
	AccountType string `json:"account_type"`
}

func InsertCustomer (customer CustomerORM, db *gorm.DB){
	if err :=  db.Create(&customer).Error;err != nil {
		log.Println("failed to insert :",err.Error())
		return
	}
	log.Println("Success insert data")
}

func GetCustomers (db *gorm.DB){
	var customer []CustomerORM
	if err := db.Preload("AccountORM").Find(&customer).Error;err != nil{
		log.Println("failed to get data :",err.Error())
		return
	}
	log.Println(customer)
}

func DeleteCustomer (id int, db *gorm.DB){
	var customer CustomerORM
	if err := db.Where(&CustomerORM{ID: id}).Delete(&customer).Error;err != nil{
		log.Println("failed to delete data :",err.Error())
		return
	}

	log.Println("Success delete data")
}

func UpdateCustomer (customer CustomerORM,id int, db *gorm.DB){
	if err := db.Model(&CustomerORM{}).Where(&CustomerORM{ID: id}).Updates(customer).Error; err != nil{
		log.Println("failed to update data :",err.Error())
		return
	}

	log.Println("Success update data")
}

