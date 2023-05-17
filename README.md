# gos
go server tools

server graceful shutdown
``` go
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
```

client
```
$ curl http://localhost:8888/async
task running
$ curl http://localhost:8888/quit
quit in a few seconds
```
server
```
2023/05/17 17:30:47 task start
2023/05/17 17:30:49 quit in 5 seconds
2023/05/17 17:30:54 quit from api
2023/05/17 17:30:54 closing db
2023/05/17 17:30:54 Server Shutdown: success
2023/05/17 17:30:57 task finish
2023/05/17 17:30:59 db closed
2023/05/17 17:30:59 Server exit
```

case 1
func ApiQuit()
在需要退出的时候调用，会在5秒后触发
net/http.Server的Shutdown

case 2
func RegisterOnShutdown(f func())
常驻程序注册退出程序
注册的函数会在Shutdown的时候被调用

case 3
func Register() => sync.WaitGroup.Add(1)
func Done()     => sync.WaitGroup.Done()
临时程序管控

final
func Wait()     => sync.WaitGroup.Wait()
case 2 和 case 3 中Register的函数若仍未完成，则都会在此等待（Done()的调用）