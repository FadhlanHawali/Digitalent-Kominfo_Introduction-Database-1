package main

import (
	"database/sql"
	"fmt"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/config"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func main(){
	cfg,err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db,err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	//database.InsertCustomer(database.Customer{
	//	FirstName:    "Fadhlan",
	//	LastName:     "Hawali",
	//	NpwpId: "id-1",
	//	Age:          10,
	//	CustomerType: "Premium",
	//	Street:       "Str",
	//	City:         "Yogya",
	//	State:        "Indo",
	//	ZipCode:      "55555",
	//	PhoneNumber:  "0812384",
	//},db)

	//database.GetCustomers(db)
	//database.DeleteCustomer(1,db)
	database.UpdateCustomer(30,2,db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}


func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",cfg.User,cfg.Password,cfg.Host,cfg.Port,cfg.DbName,cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}