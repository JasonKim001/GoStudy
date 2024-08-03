package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// config database connect string
	dsn := "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8"

	// connect database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	// execute basic example
	createTable(db)
	insertRecord(db)
	queryRecords(db)
	updateRecord(db)
	deleteRecord(db)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        age INT
    )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
}

func updateRecord(db *sql.DB) {
	query := `UPDATE users SET age = ? WHERE name = ?`

	_, err := db.Exec(query, 31, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully!")
}

func insertRecord(db *sql.DB) {
	query := `INSERT INTO users(name, age) VALUES (?, ?)`

	_, err := db.Exec(query, "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully inserted!")
}

func queryRecords(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteRecord(db *sql.DB) {
	query := `DELETE FROM users WHERE name = ?`

	_, err := db.Exec(query, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted!")
}
