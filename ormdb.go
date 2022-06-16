package plmorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/stand00d/plmorm/plmormtypes"
)

func OpenPLMORMConnection() {

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
}

func MigrateSchema(c *gorm.DB) error {

	var dbversion plmormtypes.SchemaVersion

	err := c.First(&dbversion)

	if err != nil {
		fmt.Println("Error getting version: ", err)
	}
	c.AutoMigrate(&plmormtypes.SchemaVersion{})
	typesversion := plmormtypes.GetTypesSchemaVersion()
	if dbversion.MajorVersion != typesversion.MajorVersion {
		fmt.Println("Major Schema version does not match:")
		plmormtypes.SetNewVersion(c, typesversion)
	}
	return nil
}