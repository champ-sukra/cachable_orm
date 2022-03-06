package orm_handler

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// SetupDatabase setup using database config, tables which require auto generated primary key
func SetupDatabase(dbConfig DBConfig, tables []string) (*gorm.DB, *UidGenerator) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	g := setupCustomerUid(db, tables)

	return db, g
}

func setupCustomerUid(db *gorm.DB, tables []string) *UidGenerator {
	g := NewUidGenerator()

	for _, table := range tables {
		var count int64
		result := db.Table(table).Select("MAX(id)").Scan(&count)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		u := uint32(count)
		g.SetUniqueIdForKey(table, u)
	}

	return g
}
