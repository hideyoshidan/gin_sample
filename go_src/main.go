package main

import (
    "github.com/hideyoshidan/go_bbs/infrastructure"
)

func main() {
    db,c := infrastructure.NewDB()
 
    r := infrastructure.NewRoute(db, c)
    r.Run()
}



// package main

// import "github.com/gin-gonic/gin"

// func main() {
//     r := gin.Default()
//     r.GET("/ping", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "pong",
//         })
//     })
//     r.Run() // listen and server on 0.0.0.0:8080
// }