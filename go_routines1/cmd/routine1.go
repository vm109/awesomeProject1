package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	databaseInsertion()
}

func databaseInsertion() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS documents (id INTEGER PRIMARY KEY, published integer, created_at text)")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go insertUser(db, true, time.Now(), &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines completed.")
}

func insertUser(db *sql.DB, published bool, created_at time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	var published_state int8 = 0
	var created_at_str = created_at.Format("2006-01-02")
	if published {
		published_state = 1
	}

	_, err := db.Exec("INSERT INTO documents (published, created_at) VALUES (?, ?)", published_state, created_at_str)
	if err != nil {
		log.Fatal(err)
	}
}
