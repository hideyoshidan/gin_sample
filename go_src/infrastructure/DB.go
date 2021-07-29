package infrastructure

// ブランクimportについて。https://www.pospome.work/entry/2017/01/29/171904
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBの構造体
type DB struct {
	Host string
    Username string
    Password string
    DBName string
	DBPort string
    Connection *gorm.DB
}

// DBのポインタを初期化して返す
func NewDB() (db *DB, c *Config){
	c = NewConfig()
	db = newDB(&DB {
		Host: c.DataBase.Develop.Host,
        Username: c.DataBase.Develop.Username,
        Password: c.DataBase.Develop.Password,
        DBName: c.DataBase.Develop.DBName,
		DBPort: c.DataBase.Develop.DBPort,
	})

	return db, c
}

// 新しいDB接続のポインタを返す
func newDB(d *DB) *DB {	
	// Gormで、mysqlと接続
	mysqlUrl := d.Username + ":" + d.Password
	mysqlUrl += "@tcp(" + d.Host + ":" + d.DBPort + ")/"
	mysqlUrl += d.DBName
	mysqlUrl += "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", mysqlUrl)
	db.SingularTable(true)
	// mysql://go_user:go@tcp(go_db:3306)/go_db?charset=utf8&parseTime=True&loc=Local

	// エラーの際は実行を継続できないとし、panicで処理を中断
	if err != nil {
		panic(err.Error())
	}

	d.Connection = db

	return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
    return db.Connection.Begin()
}

// DB接続
func (db *DB) Connect() *gorm.DB {
    return db.Connection
}