package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "myuser"
	dbname = "mydb"
)

type User struct {
	ID int64
	Name string
	Email string
}

func main() {

	secretPath := "db_key.txt"

	passwordBytes, err := os.ReadFile(secretPath)
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler o arquivo de secret: %v", err))
	}

	password := strings.TrimSpace(string(passwordBytes))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}


	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) not null,
		email VARCHAR(50) not null
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(fmt.Sprintf("Erro ao criar a tabela: %v", err))
	}

	insertSql := ` INSERT INTO users (name, email) VALUES ('jon calhou', 'jon@calhoun.io')`
	res, err := db.Exec(insertSql)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	querySql := `
	SELECT * FROM users WHERE id = $1
	`

	var u User
	err = db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)


}
