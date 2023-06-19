# gos
go server tools

## 【01】server graceful shutdown
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

### client
```
$ curl http://localhost:8888/async
task running
$ curl http://localhost:8888/quit
quit in a few seconds
```
### server
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

### case 1
func ApiQuit()    
在需要退出的时候调用，会在5秒后触发net/http.Server的Shutdown    
注：关于【5秒】这个时间的来源，参考：https://segmentfault.com/a/1190000043788014

### case 2
func RegisterOnShutdown(f func())    
常驻程序注册退出程序，注册的函数会在Shutdown的时候被调用

### case 3
func Register() => sync.WaitGroup.Add(1)    
func Done()     => sync.WaitGroup.Done()    
临时程序管控

### final
func Wait()     => sync.WaitGroup.Wait()    
case 2 和 case 3 中Register的函数若仍未完成，则都会在此等待（Done()的调用）

***

## 【02】timeo for time offset
服务器debug时，经常需要修改时间
为了避免频繁修改服务器本地时间影响其他服务，可以修改服务的内置时间
实现起来就是给时间加上一个offset

目前设置time offset的方式有两种：
全局式和对象式

全局式
```
timeo.SetOffset(timeo.Day*3 + timeo.Hour*12)
now := timeo.Now()          // time.Time
oss := timeo.GetOffset()    // int64
fmt.Println(now, time.Duration(oss) * time.Second)  // 2023-06-23 04:54:38 +0800 CST 84h0m0s
```

对象式
```
pos := timeo.NewOffset(timeo.Day*3 + timeo.Hour*12)
now := pos.Now()            // time.Time
oss := pos.GetOffset()      // int64
fmt.Println(now, time.Duration(oss) * time.Second)  // 2023-06-23 04:54:38 +0800 CST 84h0m0s
```
