package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price 		     string `json:"priced"`
}

type receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []item `json:"items"`
	Total        string `json:"total"`
	Points       string `json:"points"`
	Id		     string `json:"id"`
}

var idMap = make(map[string]string)

var items1 = []item {
	{
		ShortDescription: "Mountain Dew",
		Price: "12.45",
	},
	{
		ShortDescription: "Water",
		Price: "3.04",
	},
	{
		ShortDescription: "Coffee",
		Price: "10.99",
	},
}

var items2 = []item {
	{
		ShortDescription: "desc1",
		Price: "12.34",
	},
	{
		ShortDescription: "desc2",
		Price: "15.09",
	},
	{
		ShortDescription: "desc3",
		Price: "8.75",
	},
}

var receipts = []receipt{
    {   
		Retailer: "Woodman's", 
	    PurchaseDate: "2020-02-01", 
	    PurchaseTime: "13:01", 
	    Items: items1,
	    Total: "26.48",
	},
	{
		Retailer: "Walmart", 
		PurchaseDate: "2022-04-29", 
		PurchaseTime: "8:06", 
		Items: items2,
		Total: "36.18",
	},
}


func main() {
	fmt.Println("in main.go main() method")
	router := gin.Default()
	router.GET("/receipts", getAllReceipts)
	router.POST("/receipts", postReceipt)
	router.GET("/receipts/:id/points", getReceiptById)

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
 	points, err := calculatePoints(&newReceipt)

	if err != nil {
		panic(err)
	}
	newReceipt.Points = points

	// assign ID
	newReceipt.Id = uuid.New().String()

	fmt.Println(newReceipt.Id)

	// put [id]points entry into map for later retrieval
	idMap[newReceipt.Id] = newReceipt.Points
	receipts = append(receipts, newReceipt)
	response := "id: " + newReceipt.Id 
	cntx.IndentedJSON(http.StatusCreated, response)
}

func calculatePoints(r *receipt) (string, error) {
	points := 0

	if len(r.Total) != 5 {
		return "", errors.New("Incorrect number format for receipt total\n")
	}
	
	// 1 point per alphanumeric character in retailer name
	points += len(strings.ReplaceAll(r.Retailer, " ", ""))

	suffix := r.Total[3:]
	// 50 points if total amount is round dollar amount (ends in .00)
	// 25 points if total amount is a multiple of .25
	switch suffix {
		case "00":
			points += 75
		case ".25", ".50", ".75":
			points += 50	
	}

	// 5 points per 2 items on receipt
	points += (len(r.Items)/2) * 5

	// trimmed length of item description
	for _, item := range r.Items {
		if(len(strings.ReplaceAll(item.ShortDescription, " ", ""))) % 3 == 0 {
			multipliedPrice, err := strconv.ParseFloat(item.Price, 64)
			if(err != nil) {
				return "", errors.New("Inccorect format for price of " + item.ShortDescription)
			}
			multipliedPrice = math.Round(multipliedPrice * .2)

		}
	}

	


	return strconv.Itoa(points), nil
}

