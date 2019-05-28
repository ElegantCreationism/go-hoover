package service

import (
	"context"
	"flag"
	"github.com/ElegantCreationism/go-hoover/env"
	"github.com/ElegantCreationism/go-hoover/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	log.Printf(env.Settings.Address + env.Settings.Port)
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/roomba", handler.RoombaHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Addr:         env.Settings.Address + env.Settings.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ERROR STARTING SERVER :%v", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("ERROR SHUTTING DOWN SERVICE: %v", err)
	}

	log.Println("shutting down")
	os.Exit(0)

}
