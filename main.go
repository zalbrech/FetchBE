package main

import (
	// "errors"
	// "fmt"
	// "math"
	"net/http"
	// "regexp"

	// "strconv"
	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []item `json:"items"`
	Total        string `json:"total"`
	Points       string `json:"points"`
	Id           string `json:"id"`
}

var idMap = make(map[string]string)

var receipts = []receipt{}

func main() {
	router := gin.Default()
	router.GET("/receipts", getAllReceipts)
	router.GET("/receipts/:id/points", getReceiptById)
	router.POST("/receipts", postReceipt)
	
	router.Run("localhost:8080")
}

func getAllReceipts(cntx *gin.Context) {
	cntx.IndentedJSON(http.StatusOK, receipts)
}

func getReceiptById(cntx *gin.Context) {
	id := cntx.Param("id")
	cntx.IndentedJSON(http.StatusOK, idMap[id])
}

func postReceipt(cntx *gin.Context) {
	var newReceipt receipt

	if err := cntx.BindJSON(&newReceipt); err != nil {
		panic(err)
	}
	// calculate points
	points, err, pointsSlice := calculatePoints(&newReceipt)

	if err != nil || pointsSlice == nil {
		panic(err)
	}
	newReceipt.Points = points

	// assign ID
	newReceipt.Id = uuid.New().String()

	// put [id]points entry into map for later retrieval
	idMap[newReceipt.Id] = newReceipt.Points

	// add newReceipt to 'database'
	receipts = append(receipts, newReceipt)

	response := "id: " + newReceipt.Id
	cntx.IndentedJSON(http.StatusCreated, response)
}
