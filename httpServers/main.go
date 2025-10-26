package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)


func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	
	srv := &http.Server{
		Addr:                         "localhost:8080",
		Handler:                      log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
		}
	}
}