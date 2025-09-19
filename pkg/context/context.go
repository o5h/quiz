package context

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/o5h/config"
)

var (
	ctx           context.Context
	cancel        context.CancelFunc
	shutdownFuncs []ShutdownFunc
)

type ShutdownFunc func(context.Context) error

func Init(version, date string) {
	log.Println("Quiz Server", "version:", version, "build date:", date)

	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	go func() {
		<-ctx.Done() // Wait for interrupt signal
		shutdownTimeout := time.Duration(config.Get("server.shutdown_timeout", 5)) * time.Second
		log.Println("Received interrupt signal, shutting down with timeout:", shutdownTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := onShutdown(ctx); err != nil {
			log.Fatalln("Error during shutdown:", err)
		}
		os.Exit(0)
	}()

}

func Shutdown() {
	log.Println("Context: closing")
	cancel()
}

func Get() context.Context {
	return ctx
}

func RegisterShutdown(fn ShutdownFunc) {
	shutdownFuncs = append(shutdownFuncs, fn)
}

// onShutdown calls all registered shutdown functions
func onShutdown(ctx context.Context) error {
	for _, fn := range shutdownFuncs {
		if err := fn(ctx); err != nil {
			return err
		}
	}
	return nil
}
