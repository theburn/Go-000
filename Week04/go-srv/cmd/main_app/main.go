package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const listenAddr = ":8081"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word!")
}

func listHttpServ(srv *http.Server) error {
	fmt.Println("listening at " + listenAddr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func shutdownClean(ctx context.Context, srv *http.Server) error {
	select {
	case <-ctx.Done():
		fmt.Println("close timeout, immediately closed!")
	default:
		_doClean()
	}
	return srv.Shutdown(ctx)
}

func _doClean() {
	fmt.Println("do something clean... !")
}

func exitSignalHandler(srv *http.Server, exit chan os.Signal) error {
	<-exit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//time.Sleep(16 * time.Second)

	return shutdownClean(ctx, srv)
}

func main() {

	g := new(errgroup.Group)

	srv := http.Server{
		Addr:    listenAddr,
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/", helloHandler)

	// register signal
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		return exitSignalHandler(&srv, exit)
	})

	g.Go(func() error {
		return listHttpServ(&srv)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Shutdown with Error: " + err.Error())
	} else {
		fmt.Println("Shutdown finished!")
	}

}
