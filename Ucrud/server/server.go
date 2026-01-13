package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"banco/db"
)

type User struct {
	ID        uint32 `json:"id"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error to read body", http.StatusBadRequest)
		return
	}

	var user User
	if err = json.Unmarshal(responseBody, &user); err != nil {
		http.Error(w, "error to unmarshal", http.StatusBadRequest)
		fmt.Println("Error to unmarshal:", err) 
		return
	}

	db, err := db.InitDB()
	if err != nil {
		w.Write([]byte("error to conect database"))
		fmt.Println("Error to connect database:", err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (age, first_name, last_name, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		w.Write([]byte("Error to prepare statement"))
		log.Printf("Error: %v", err)
		return
	}
	defer statement.Close()
	
	insert, err := statement.Exec(user.Age, user.FirstName, user.LastName, user.Email)
	if err != nil {
		w.Write([]byte("Error to insert user"))
		log.Printf("Error: %v", err)
		return
	}

	rows, err := insert.RowsAffected()
	if err != nil {
		w.Write([]byte("Error to get affected rows"))
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created successfully, %d rows affected", rows)))
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.InitDB()
	if err != nil {
		w.Write([]byte("error to connect database"))
		fmt.Println("Error to connect database:", err)
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Error to query users"))
		log.Printf("Error: %v", err)
		return
	}
	defer row.Close()

	var users []User
	for row.Next() {
		var user User
		if err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.LastName); err != nil {
			w.Write([]byte("Error to scan user"))
			log.Printf("Error: %v", err)
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to encode users"))
		log.Printf("Error: %v", err)
		return
	}

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("error to convert id"))
		return
	}

	db, err := db.InitDB()
	if err != nil {
		w.Write([]byte("error to connect database"))
		fmt.Println("Error to connect database:", err)
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = $1", ID)
	if err != nil {
		w.Write([]byte("Error to query user"))
		log.Printf("Error: %v", err)
		return
	}

	var user User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			w.Write([]byte("Error to scan user"))
			log.Printf("Error: %v", err)
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error to encode user"))
		log.Printf("Error: %v", err)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("error to convert id"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error to read body", http.StatusBadRequest)
		return
	}

	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		http.Error(w, "error to unmarshal", http.StatusBadRequest)
		fmt.Println("Error to unmarshal:", err)
		return
	}

	db, err := db.InitDB()
	if err != nil {
		w.Write([]byte("error to connect database"))
		fmt.Println("Error to connect database:", err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET age=$1, first_name=$2, last_name=$3, email=$4 WHERE id=$5"); if err != nil {
		w.Write([]byte("Error to prepare statement"))
		log.Printf("Error: %v", err)
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Age, user.FirstName, user.LastName, user.Email, ID); err != nil {
		w.Write([]byte("Error to update user"))
		log.Printf("Error: %v", err)
		return
	}
}