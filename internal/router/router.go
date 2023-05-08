package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"tde/internal/microservices/logger"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const gracefulShutdownTimeout = 15 * time.Second

var servers = []*http.Server{}

var log = logger.NewLogger("Router")

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not found. Please return homepage.")
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func StartRouter(baseURL string, endpointRegisterer func(r *mux.Router)) {
	r := mux.NewRouter()
	endpointRegisterer(r)
	// r.Use(middleware.MWAuthorization)
	r.Use(mux.CORSMethodMiddleware(r))

	server := &http.Server{
		Addr: baseURL,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Calling ListenAndServe")
		if err := server.ListenAndServe(); err != nil {
			log.Println(errors.Wrap(err, "http.Server returned an error from ListendAndServe call"))
		}
	}()

	servers = append(servers, server)
}

func StartTLSRouter(baseURL string, endpointRegisterer func(r *mux.Router)) {
	r := mux.NewRouter()
	endpointRegisterer(r)
	// r.Use(middleware.MWAuthorization)
	r.Use(mux.CORSMethodMiddleware(r))

	server := &http.Server{
		Addr: baseURL,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Calling ListenAndServeTLS")
		if err := server.ListenAndServeTLS(publicCertPath, privateCertPath); err != nil {
			log.Println(errors.Wrap(err, "http.Server returned an error from ListenAndServeTLS call"))
		}
	}()

	servers = append(servers, server)
}

func Wait() {
	sigInterruptChannel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sigInterruptChannel, os.Interrupt)

	// Block until we receive our signal.
	<-sigInterruptChannel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()

	for _, server := range servers {
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		log.Printf("Sending shutdown signal to one of the servers, grace period is '%s'\n", gracefulShutdownTimeout.String())
		go server.Shutdown(ctx)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	<-ctx.Done()
	log.Println("All servers are closed")
}
