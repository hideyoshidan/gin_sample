package database

import (
    "github.com/jinzhu/gorm"
)

// DBの構造体
type DB interface {
    Begin() *gorm.DB
    Connect() *gorm.DB
}