package db

import (
	"log"

	"github.com/go-sql-driver/mysql"
	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDb struct {
	db *gorm.DB
}

func Open(c mysql.Config, opts ...gorm.Option) *MysqlDb {
	db, err := gorm.Open(mysqlGorm.Open(c.FormatDSN()), opts...)
	if err != nil {
		log.Fatalf("failed to open mysql conn: %v", err)
	}

	return &MysqlDb{
		db: db,
	}
}

func (d *MysqlDb) GetDB() *gorm.DB {
	return d.db
}

func (d *MysqlDb) SetDebug(e bool) {
	if e {
		d.db = d.db.Debug()
	}
}
