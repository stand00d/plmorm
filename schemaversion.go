package plmorm

import (
	"time"

	"github.com/jinzhu/gorm"
)

type SchemaVersion struct {
	gorm.Model
	MajorVersion int `sql:"type int"`
	MinorVersion int `sql:"type int"`
}

func GetTypesSchemaVersion() SchemaVersion {
	var v SchemaVersion
	v.CreatedAt = time.Now()
	v.MajorVersion = 0
	v.MinorVersion = 1
	return v
}
func SetNewVersion(g *gorm.DB, v SchemaVersion) error {

	exists := g.NewRecord(v)
	if !exists {
		x := v
		x.CreatedAt = time.Now()
		x.UpdatedAt = time.Now()
		g.Create(x)
	} else {
		x := v
		x.UpdatedAt = time.Now()
		g.Update(&x)
	}
	return nil
}
