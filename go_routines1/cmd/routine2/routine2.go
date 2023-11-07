package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Document struct {
	Id        int    `db:"id"`
	Published int    `db:"published"`
	CreatedAt string `db:"created_at"`
}

func main() {
	fmt.Println("---- routines scenario 2 ----")

	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	insertDocumentsIntoDatabase(db, 100)
	documents := getDocuments(db)

	fmt.Println(len(documents))
	deleteDocumentBefore(db, documents, "2023-10-29")
	fmt.Println(len(getDocuments(db)))
}

func insertDocumentsIntoDatabase(db *sql.DB, records int) {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS documents (id INTEGER PRIMARY KEY, published integer, created_at text)")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < records; i++ {
		insertDocument(db, true, time.Now())
	}

	fmt.Println("added documents successfull")
}

func insertDocument(db *sql.DB, published bool, created_at time.Time) {
	var published_state int = 0
	var created_at_str = created_at.Format("2006-01-02")
	if published {
		published_state = 1
	}

	_, err := db.Exec("INSERT INTO documents (published, created_at) VALUES (?, ?)", published_state, created_at_str)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteDocumentBefore(db *sql.DB, documents []Document, date string) {
	// Spin 5 worker goroutines
	// workers Read documents from database
	// delete those documents if matches with date string
	// once done spin 5 more goroutines

	numOfWorkers := 5
	var wg sync.WaitGroup
	initial := 0
	interval := int(math.Round(float64(len(documents)) / float64(5)))
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go deleteWorker(db, documents[initial:initial+interval], date, &wg)
		initial = initial + interval
	}

	wg.Wait()
	fmt.Printf("successfully deleted all documents which match the date %s", date)
}

func deleteWorker(db *sql.DB, documents []Document, date string, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(documents) > 0 {
		for _, document := range documents {
			if strings.Contains(document.CreatedAt, date) {
				_, err := db.Exec("DELETE FROM documents where id=$1", document.Id)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func getDocuments(db *sql.DB) (documents []Document) {

	rows, err := db.Query("SELECT id, published, created_at FROM documents limit 100")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var doc Document
		err = rows.Scan(&doc.Id, &doc.Published, &doc.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}

		documents = append(documents, doc)
	}

	return documents
}
