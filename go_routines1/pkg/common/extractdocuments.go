package common

import (
	"fmt"
	"go_routines1/pkg/model"
	"sync"
)

type ExtractedDocument struct {
	totalDocuments int
	website        string
}

func ExtractDocuments(documents []model.Document, published bool, websites []string) {
	websiteChannel := make(chan string, len(websites))
	aggregatorChannel := make(chan ExtractedDocument, len(websites))
	var websitesWaitGroup sync.WaitGroup

	totalWebsites := len(websites)
	documentsPerWebsite := len(documents) / totalWebsites

	// Start goroutines first
	for i := 0; i < totalWebsites; i++ {
		websitesWaitGroup.Add(1)

		go websiteWorker(&websitesWaitGroup, websiteChannel, documents[i*documentsPerWebsite:(i+1)*documentsPerWebsite], aggregatorChannel)
	}

	// Send values to the channel after starting goroutines

	for _, website := range websites {
		websiteChannel <- website
	}
	close(websiteChannel)
	websitesWaitGroup.Wait()

	close(aggregatorChannel)
	// Read from aggregatorChannel
	for result := range aggregatorChannel {
		fmt.Println(result)
	}
}

func websiteWorker(wg *sync.WaitGroup, websiteChannel <-chan string, documents []model.Document, aggregatorChannel chan<- ExtractedDocument) {
	defer wg.Done()
	website := <-websiteChannel
	total := 0
	for _, document := range documents {
		if document.WebsiteId == website {
			total++
		}
	}
	result := ExtractedDocument{
		totalDocuments: total,
		website:        website,
	}
	aggregatorChannel <- result
}
