package app

import (
	"context"
	"log"
	"time"
)

// serverApp является абстрактным приложением.
type serverApp struct {
	ctx context.Context
}

// NewServerApp - конструктор приложения.
func NewServerApp(ctx context.Context) *serverApp {
	return &serverApp{ctx: ctx}
}

// Run - основной цикл программы.
func (a *serverApp) Run() error {
	go a.exampleWorker()

	<-a.ctx.Done()
	log.Println("[DEBUG] application terminated")
	a.Shutdown()
	return nil
}

// Shutdown - выключает сервер, закрывая канал.
func (a *serverApp) Shutdown() {
	log.Println("[DEBUG] serverApp shutdown successfully")
}

// exampleWorker выводит сообщение о своей работе на экран.
func (a *serverApp) exampleWorker() {
	for {
		select {
		case <-time.After(time.Second):
			log.Println("[INFO] work in progress...")
		case <-a.ctx.Done():
			break
		}
	}
}
