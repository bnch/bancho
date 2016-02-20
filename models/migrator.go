package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sort"
	// Hello golint pls dont break balls kthx bye
	_ "github.com/Go-SQL-Driver/MySQL"
)

// Temporary until we have a configuration system.
const (
	DBUser = "root"
	DBPass = ""
	DBName = "bancho"
	DBHost = "tcp(localhost:3306)"
	// Database type
	DB = "mysql"
)

// Migrate automatically migrates the passed database to the latest version.
func Migrate(db gorm.DB) error {
	fmt.Println("==> migrating database...")
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DBVer{})
	v := DBVer{}
	keys := []uint64{}
	for k := range migrations {
		keys = append(keys, k)
	}
	sort.Sort(uint64arr(keys))
	for _, k := range keys {
		if k > v.Version {
			migrations[k](db)
		}
		v.Version = k
	}
	db.Save(&v)
	fmt.Println("==> database migrated!")
	return nil
}

// CreateDB creates an instance of a Gorm database.
func CreateDB() (gorm.DB, error) {
	db, err := gorm.Open(DB, DBUser+":"+DBPass+"@"+DBHost+"/"+DBName+"?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	return db, err
}
