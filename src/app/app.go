package app

import (
	"context"
	"log"
	"time"
)

type serverApp struct {
	terminated chan struct{}
}

func NewServerApp() *serverApp {
	return &serverApp{
		terminated: make(chan struct{}),
	}
}

// Wait() - ожидает завершение работы приложения.
func (a *serverApp) Wait() {
	<-a.terminated
}

// Run() - основной цикл программы.
func (a *serverApp) Run(ctx context.Context) error {
	for {
		select {
		case <-a.terminated:
			log.Println("[DEBUG] serverApp terminated")
			return nil
		case <-ctx.Done():
			log.Println("[DEBUG] application terminated")
			return nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	a.Shutdown()

	return nil
}

// Shutdown() - выключает сервер, закрывая канал.
func (a *serverApp) Shutdown() {
	close(a.terminated)
}
