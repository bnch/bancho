package models

import (
	"fmt"
	"github.com/bnch/bancho/conf"
	"github.com/jinzhu/gorm"
	"sort"
	// Hello golint pls dont break balls kthx bye
	_ "github.com/Go-SQL-Driver/MySQL"
)

var cachedDB *gorm.DB

// Migrate automatically migrates the passed database to the latest version.
func Migrate(db gorm.DB) error {
	fmt.Println("==> migrating database...")
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DBVer{})
	v := DBVer{}
	db.First(&v)
	keys := []uint64{}
	for k := range migrations {
		keys = append(keys, k)
	}
	sort.Sort(uint64arr(keys))
	for _, k := range keys {
		if k > v.Version {
			migrations[k](db)
			v.Version = k
		}
	}
	db.Save(&v)
	fmt.Println("==> database migrated!")
	return nil
}

// CreateDB creates an instance of a Gorm database.
func CreateDB() (gorm.DB, error) {
	if cachedDB != nil {
		return *cachedDB, nil
	}
	c, err := conf.Get()
	if err != nil {
		return gorm.DB{}, err
	}
	db, err := gorm.Open(
		c.SQLInfo.DBType,
		c.SQLInfo.User+":"+c.SQLInfo.Pass+"@"+c.SQLInfo.Host+"/"+c.SQLInfo.Name+"?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		cachedDB = &db
	}
	return db, err
}
