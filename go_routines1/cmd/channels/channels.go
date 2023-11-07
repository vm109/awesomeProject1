package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("start channel communication")
	createDocumentsForWebsites(204, 4)
}

type document struct {
	id        string
	website   string
	published bool
	createdat *time.Time
}

func createDocumentsForWebsites(totalDocuments int, totalWebsites int) {
	documentChannel := make(chan document, totalDocuments)

	var wg sync.WaitGroup
	for i := 0; i < totalWebsites; i++ {
		website := fmt.Sprintf("website%d", i)
		documentShare := int(math.Round(float64(totalDocuments) / float64(totalWebsites)))
		wg.Add(1)
		go generateRandomDocuments(&wg, documentChannel, website, documentShare)
	}

	wg.Wait()
	close(documentChannel)

	n := 0
	for doc := range documentChannel {
		fmt.Println(doc, n)
		n++
	}
}

func generateRandomDocuments(wg *sync.WaitGroup, docChannel chan<- document, website string, documentCount int) {
	defer wg.Done()
	fmt.Println("-------")
	fmt.Println(documentCount)
	fmt.Println("-------")
	for i := 0; i < documentCount; i++ {
		id := generateRandomId()
		t := generateRandomTime()
		published := generateRandomPublished()
		doc := document{
			id:        id,
			createdat: t,
			website:   website,
			published: published,
		}

		docChannel <- doc
	}
}

func generateRandomPublished() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func generateRandomId() string {
	return fmt.Sprintf("ID_%d", rand.Intn(1000))
}

func generateRandomTime() *time.Time {
	t := time.Now().Add(-time.Duration(rand.Intn(365*24*60*60)) * time.Second)
	return &t
}
