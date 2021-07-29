package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/hideyoshidan/go_bbs/interfaces/controllers"
)

// ルーティングの構造体
type Routing struct {
	DB *DB
	Gin *gin.Engine
	Port string
}

// Routingを初期化したポインタを返す
func NewRoute(db *DB, c *Config) *Routing {
	r := &Routing {
		DB: db,
		Gin: gin.Default(),
		Port: c.Route.Port,
	}

	r.setRouting()
	return r
}

// ルーティングをセット
func (r *Routing) setRouting (){
	sampleController := controllers.NewSampleController(r.DB)
	r.Gin.GET("sample/:id", func (c *gin.Context) { sampleController.Get(c) })
}

// 実行
func (r *Routing) Run() {
    // r.Gin.Run(r.Port) PORTを指定するとなぜか使えない
	r.Gin.Run()
}