package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data any `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err:= json.Marshal(resp)
	if err != nil {
		slog.Error("error when trying do marshal", "error", err)
		sendJSON(w, Response{Error:"something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil{
		slog.Error("error to send response", "error", err)
		return
	}

}

type User struct {
	ID int64 `json:",string"` // para numeros enormes, o discord faz isso
	Username string
	Role string
	Password Password `json:"-"` //ignorar o campo, pra ele n vir no body
}

type Password string

func (p Password) LogValue() slog.Value{
	return slog.StringValue("[REDACTED]")
}

func main() {
	p := Password("hsgdhsdg")
	slog.Info("Password", "p", p)
	
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     nil,
		ReplaceAttr: nil,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))//log to json
	slog.SetDefault(logger) // set all log default
	logger = slog.With(slog.Group("app_info", slog.String("version", "1.0.0")))
	slog.Info("iniciating service", "time", time.Now(), "version", "1.0.0")
	logger.LogAttrs(context.Background(), slog.LevelInfo, "http request",
		slog.Group("http_data", // output explained in terminal
			slog.String("method", http.MethodDelete),
		) ,
		slog.Duration("time_tanken", time.Second),
		slog.Int("status", http.StatusOK),
		)

	//slog.Warn("test") -> another type of log
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {
			ID: 1,
			Username:"admin",
			Role: "admin",
			Password: "admin",
		},
	} //O id e o objeto, o usuario

	r.Group(func(r chi.Router)  {
		r.Use(jsonMiddleware)//função use é para middleware
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db)) //regex para garantir que seja numero
		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler { //é boa pratica o content-type da requisição retornar como json
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r) // chama o proximo handle pelo o q eu entendi, ver a linha 36
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64) //estamos ignorando um erro, ja q o regex n permite ele ter, mas mesmoa ssim ele pode ocorrer se o numero for overflow. NÃO ´É BOA PRATICA

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
	r.Body = http.MaxBytesReader(w, r.Body, 100000)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr){
			sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
			return
		}

		slog.Error("fail to read user", "error", err)
		sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
		return
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	db[user.ID] = user

	w.WriteHeader(http.StatusCreated)
  }
}