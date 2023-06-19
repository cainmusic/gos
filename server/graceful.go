package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
graceful
一、使用Shutdown和RegisterOnShutdown结束主服务
二、使用WaitGroup处理临时的goroutine以最终结束全部服务
*/

type MyServer struct {
	srv      *http.Server
	fHandler bool
}

type WaitGroup struct {
	wg sync.WaitGroup
}

var msrv *MyServer
var wg *WaitGroup
var once sync.Once

var apiQuit = make(chan bool)
var quit = make(chan os.Signal)

func init() {
	once.Do(func() {
		msrv = &MyServer{
			srv: &http.Server{
				Addr: ":8888",
			},
			fHandler: false,
		}
		wg = &WaitGroup{}
	})
}

// MyServer

func ServerSetAddr(addr string) {
	msrv.srv.Addr = addr
}

func ServerSetHandler(handler http.Handler) {
	msrv.srv.Handler = handler
	msrv.fHandler = true
}

func Run() {
	if !msrv.fHandler {
		log.Fatal("no handler set")
	}

	go func() {
		log.Println("listening", msrv.srv.Addr)
		if err := msrv.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signal.Notify(quit, os.Interrupt)

	WaitingForSignal()

	Shutting()

	Wait()

	log.Println("Server exit")
}

func ApiQuit() {
	log.Println("quit in 5 seconds")
	time.AfterFunc(5*time.Second, func() { apiQuit <- true })
}

func WaitingForSignal() {
	select {
	case <-quit:
		log.Println("quit from os.Signal")
	case <-apiQuit:
		log.Println("quit from api")
	}
}

func Shutting() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := msrv.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Shutdown: success")
}

func RegisterOnShutdown(f func()) {
	Register()
	msrv.srv.RegisterOnShutdown(func() {
		defer Done()
		f() // 这个函数应该是阻塞的
	})
}

// WaitGroup

func Wait() {
	wg.wg.Wait()
}

func Register() {
	wg.wg.Add(1)
}

func Done() {
	wg.wg.Done()
}
