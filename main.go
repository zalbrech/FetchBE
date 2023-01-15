package main

import (
	"fmt"
)

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price 		     string `json:"priced"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []item `json:"items"`
	Total string `json:"total"`
}


func main() {
	// router := gin.Default()
	fmt.Println("in main.go main() method")
}

// func getAlbums(c *gin.Context) {
//     c.IndentedJSON(http.StatusOK, albums)
// }

