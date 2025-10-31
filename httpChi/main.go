package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID int64 `json:",string"` // para numeros enormes, o discord faz isso
	Username string
	Role string
	Password string `json:"-"` //ignorar o campo, pra ele n vir no body
}

func main() {
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
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		data, err := json.Marshal(user)
		if err !=nil {
			http.Error(w, "Something going wrong", http.StatusInternalServerError)
			return 
		}

		_, _ = w.Write(data)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
	r.Body = http.MaxBytesReader(w, r.Body, 100000)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr){
			http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
			return
		}
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		http.Error(w, "invalid body", http.StatusUnprocessableEntity)
		return
	}

	db[user.ID] = user

	w.WriteHeader(http.StatusCreated)
  }
}