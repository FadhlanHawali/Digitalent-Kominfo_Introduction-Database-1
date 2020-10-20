package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Customer struct {
	ID int `json:"id"`
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
}


func InsertCustomer (customer Customer, db *sql.DB){
	_, err := db.Exec("insert into customers(first_name,last_name,npwp_id,age,customer_type,street,city,state,zip_code,phone_number) values  (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		customer.FirstName,
		customer.LastName,
		customer.NpwpId,
		customer.Age,
		customer.CustomerType,
		customer.Street,
		customer.City,
		customer.State,
		customer.ZipCode,
		customer.PhoneNumber)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("insert success!")
}

func GetCustomers (db *sql.DB){
	rows, err := db.Query("select * from customers")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Customer

	for rows.Next() {
		var each = Customer{}
		var err = rows.Scan(
			&each.ID,
			&each.FirstName,
			&each.LastName,
			&each.NpwpId,
			&each.Age,
			&each.CustomerType,
			&each.Street,
			&each.City,
			&each.State,
			&each.ZipCode,
			&each.PhoneNumber)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	log.Println(result)
}

func DeleteCustomer (id int, db *sql.DB){
	_, err := db.Exec("delete from customers where customer_id = ?", id)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func UpdateCustomer (age int, id int, db *sql.DB){
	_, err := db.Exec("update customers set age = ? where customer_id = ?", age, id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("update success!")
}

