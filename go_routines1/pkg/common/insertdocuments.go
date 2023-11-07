package common

import (
	"fmt"
	"go_routines1/pkg/model"
	"math/rand"
	"time"
)

func InsertRandomDocuments(numberofdocuments, numberofwebsites int) []model.Document {
	var documents []model.Document = make([]model.Document, 0)
	for i := 0; i < numberofwebsites; i++ {
		documents = generateRandomDocuments(documents, numberofdocuments/numberofwebsites, fmt.Sprintf("website_%d", i))
	}
	return documents
}

func generateRandomDocuments(documents []model.Document, numberofdocuments int, websiteId string) []model.Document {
	for i := 0; i < numberofdocuments; i++ {
		document := model.Document{
			Id:        generateRandomId(),
			Published: generateRandomPublished(),
			CreatedAt: generateRandomTime(),
			WebsiteId: websiteId,
		}
		documents = append(documents, document)
	}
	return documents
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
