package httpserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

type HttpServer interface {
	Start() error
	Stop(ctx context.Context) error
	RunGracefuly() error
}

type httpServer struct {
	opt    Option
	server *http.Server
}

type Option struct {
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

func NewHttpServer(opt Option, h http.Handler) HttpServer {
	server := &http.Server{
		ReadTimeout:       opt.ReadTimeout,
		WriteTimeout:      opt.WriteTimeout,
		IdleTimeout:       opt.IdleTimeout,
		ReadHeaderTimeout: opt.ReadHeaderTimeout,
		Handler:           h,
	}

	return &httpServer{
		opt:    opt,
		server: server,
	}
}

func (s *httpServer) Start() error {
	// add sysinfo
	buildInfo, _ := debug.ReadBuildInfo()

	l, err := net.Listen("tcp", s.opt.Port)
	if err != nil {
		log.Fatalf("failed start http server. err=%s", err)
		return err
	}
	log.Printf("server running on port %s", s.opt.Port)
	log.Printf("pid: %d", os.Getpid())
	log.Printf("go_version: %s", buildInfo.GoVersion)
	return s.server.Serve(l)
}

func (s *httpServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *httpServer) RunGracefuly() error {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGINT,
	)

	go s.Start()

	sig := <-signalChannel // graceful shutdown
	log.Printf("receiving terminate signal: %s", sig)
	signal.Stop(signalChannel)
	close(signalChannel)

	if err := s.Stop(context.Background()); err != nil {
		log.Printf("failed to stop http server: %v", err)
		return err
	}
	log.Printf("http server stopped!")
	return nil
}
