package main

import (
	"log"
	"time"

	"github.com/cainmusic/gos/server"
	"github.com/cainmusic/gos/timeo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//InitDB()
	InitRouter(router)

	router.GET("/async", func(c *gin.Context) {
		log.Println("task start")
		go func() {
			server.Register()
			defer server.Done()
			time.Sleep(10 * time.Second)
			log.Println("task finish")
		}()
		c.String(200, "task running")
	})

	router.GET("/quit", func(c *gin.Context) {
		server.ApiQuit()
		c.String(200, "quit in a few seconds")
	})

	server.ServerSetHandler(router)
	server.ServerSetAddr(":8080")
	server.Run()
}

func InitDB() {
	server.RegisterOnShutdown(func() {
		// close db
		log.Println("closing db")
		time.Sleep(5 * time.Second)
		log.Println("db closed")
	})
}

func InitRouter(r *gin.Engine) {
	initTimeoGlobal(r)
	initTimeoForMe(r)
}

func initTimeoGlobal(r *gin.Engine) {
	r.GET("/setoffset1day", func(c *gin.Context) {
		timeo.SetOffset(timeo.Day)
		c.String(200, "set offset 1 day")
	})
	r.GET("/setoffset-1day", func(c *gin.Context) {
		timeo.SetOffset(-timeo.Day)
		c.String(200, "set offset -1 day")
	})
	r.GET("/gettime", func(c *gin.Context) {
		now := timeo.Now()
		c.String(200, now.String())
	})
}

func initTimeoForMe(r *gin.Engine) {
	r.GET("/setoffset1day_forme", func(c *gin.Context) {
		mytime := timeo.NewOffset(timeo.Day)
		c.String(200, "set offset 1 day : " + mytime.Now().String())
	})
	r.GET("/setoffset-1day_forme", func(c *gin.Context) {
		mytime := timeo.NewOffset(-timeo.Day)
		c.String(200, "set offset -1 day : " + mytime.Now().String())
	})
}
