package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "host=baas-rds-dev-725b87755a61c35c.elb.us-west-2.amazonaws.com user=dev_admin password=postgre&308 dbname=baas_dev_db port=1025 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error While Connecting to database")
	} else {
		fmt.Println("Database Connected")
	}

	// db.AutoMigrate()
	return db
}
