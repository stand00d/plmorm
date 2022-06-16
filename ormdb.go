package plmorm

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func OpenPLMORMConnection() *gorm.DB {

	dbConn, err := gorm.Open("mysql", "gormdba:gormdba@tcp(127.0.0.1:3306)/gormdb?charset=utf8&parseTime=true")

	if err != nil {
		fmt.Println(err)
	}

	// init
	dbConn.DB()
	dbConn.DB().Ping()
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(100)

	err = MigrateSchema(dbConn)
	if err != nil {
		fmt.Println(err)
	}
	return dbConn
}

func MigrateSchema(c *gorm.DB) error {

	var dbversion SchemaVersion

	err := c.First(&dbversion)

	if err != nil {
		fmt.Println("Error getting version: ", err)
	}
	c.AutoMigrate(&SchemaVersion{})
	typesversion := GetTypesSchemaVersion()
	if dbversion.MajorVersion != typesversion.MajorVersion {
		fmt.Println("Major Schema version does not match:")
		SetNewVersion(c, typesversion)
	}
	return nil
}
