package main

import (
	"log"
	"time"

	"github.com/cainmusic/gos/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	InitDB()

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
