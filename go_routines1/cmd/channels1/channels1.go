package main

import (
	"fmt"
	"go_routines1/pkg/common"
)

func main() {
	fmt.Println("starting golang channel communications!")

	published := true
	websites := []string{"website_0", "website_1", "website_2", "website_3"}
	documents := common.InsertRandomDocuments(100, 4)
	common.ExtractDocuments(documents, published, websites)
}
