package database

import (
    "github.com/jinzhu/gorm"
)

// DBのRepository
type DBRepository struct {
    DB DB
}

// トランザクション
func (db *DBRepository) Begin() *gorm.DB {
    return db.DB.Begin()
}

// DB接続
func (db *DBRepository) Connect() *gorm.DB {
    return db.DB.Connect()
}