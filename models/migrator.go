package models

import (
	"fmt"
	"sort"

	"github.com/bnch/bancho/conf"
	"github.com/jinzhu/gorm"
	// Hello golint pls dont break balls kthx bye
	_ "github.com/Go-SQL-Driver/MySQL"
)

var cachedDB *gorm.DB

// Migrate automatically migrates the passed database to the latest version.
func Migrate(db *gorm.DB) error {
	fmt.Println("==> migrating database...")

	c, err := conf.Get()
	if err != nil {
		return err
	}

	// In gorm 1.0, they updated the way they handle CamelCase to snake_case. Thus, DBName is no more d_b_name.
	// Which makes sense, but has broken absolutely everything.
	// So we are renaming d_b_vers to db_vers if d_b_vers still exists.
	db.Exec(`SELECT Count(*)
INTO @exists
FROM information_schema.tables 
WHERE table_schema = ?
    AND table_type = 'BASE TABLE'
    AND table_name = 'd_b_vers';

SET @query = If(@exists>0,
    'RENAME TABLE d_b_vers TO db_vers',
    'SELECT \'nothing to rename\' status');

PREPARE stmt FROM @query; 

EXECUTE stmt;`, c.SQLInfo.Name)

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
func CreateDB() (*gorm.DB, error) {
	if cachedDB != nil {
		return cachedDB, nil
	}
	c, err := conf.Get()
	if err != nil {
		return &gorm.DB{}, err
	}
	db, err := gorm.Open(
		c.SQLInfo.DBType,
		c.SQLInfo.User+":"+c.SQLInfo.Pass+"@"+c.SQLInfo.Host+"/"+c.SQLInfo.Name+"?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		cachedDB = db
	}
	return db, err
}
