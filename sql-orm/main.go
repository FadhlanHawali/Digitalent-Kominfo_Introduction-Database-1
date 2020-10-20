package main

import (
	"fmt"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/config"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-orm/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main(){
	cfg,err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db,err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}
	//
	//database.InsertCustomer(database.CustomerORM{
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
	//database.UpdateCustomer(database.CustomerORM{PhoneNumber: "0812314"},2,db)
	//
	//database.InsertAccount(database.AccountORM{
	//	Balance:         10000,
	//	AccountType:     "Deposit",
	//},2,db)
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

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&database.CustomerORM{},
		&database.AccountORM{},
		)

	log.Println("db successfully connected")

	return db, nil
}
