package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/logutils"
	"github.com/jessevdk/go-flags"

	"github.com/dbzer0/go-template/src/app"
)

var (
	revision = "unknown"
	version  = "unknown"
)

func main() {
	fmt.Printf("PROJECTNAME %s (%s)\n", version, revision)

	type options struct {
		ConfigFile string `long:"config" env:"CONFIG" description:"config file"`
		Dbg        bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
	}
	var opts options

	// парсинг опций
	p := flags.NewParser(&opts, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	setupLog(opts.Dbg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ловим сигнал для graceful termination
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Print("[WARN] interrupt signal")
		cancel()
	}()

	run(ctx)

	log.Printf("[INFO] process terminated")
}

// setupLog настраивает уровни логирования и вывод логгера в os.Stdout.
func setupLog(dbg bool) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stdout,
	}

	log.SetFlags(log.Ldate | log.Ltime)

	if dbg {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		filter.MinLevel = logutils.LogLevel("DEBUG")
	}

	log.SetOutput(filter)
}

// run запускает основной цикл программы, стартующий все остальные приложения.
func run(ctx context.Context) {
	serverApp := app.NewServerApp(ctx)

	// реализуем выключение по context cancellation
	go func() {
		<-ctx.Done()

		// имплементировать операции выключения можно здесь

		log.Printf("[DEBUG] gracefull shutdown complete!")
	}()

	serverApp.Run()
}
