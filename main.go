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

var items1 = []item{
	{
		ShortDescription: "Mountain Dew",
		Price:            "12.45",
	},
	{
		ShortDescription: "Water",
		Price:            "3.04",
	},
	{
		ShortDescription: "Coffee",
		Price:            "10.99",
	},
}

var items2 = []item{
	{
		ShortDescription: "desc1",
		Price:            "12.34",
	},
	{
		ShortDescription: "desc2",
		Price:            "15.09",
	},
	{
		ShortDescription: "desc3",
		Price:            "8.75",
	},
}

var receipts = []receipt{
	{
		Retailer:     "Woodman's",
		PurchaseDate: "2020-02-01",
		PurchaseTime: "13:01",
		Items:        items1,
		Total:        "26.48",
	},
	{
		Retailer:     "Walmart",
		PurchaseDate: "2022-04-29",
		PurchaseTime: "8:06",
		Items:        items2,
		Total:        "36.18",
	},
}

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

// func calculatePoints(r *receipt) (string, error, []int) {
// 	retailerPoints, suffixPoints, itemsLengthPoints, itemDescriptionPoints, dayPoints, timePoints := 0, 0, 0, 0, 0, 0
// 	points := 0

// 	decRegex, _ := regexp.Compile(`^\d+\.?\d\d$`)
// 	alphaNumericRegex, _ := regexp.Compile(`[^a-zA-Z0-9 ]+`)

// 	// 1 point per alphanumeric character in retailer name
// 	formattedRetailName := formatString(alphaNumericRegex, r.Retailer)
// 	retailerPoints += len(strings.ReplaceAll(formattedRetailName, " ", ""))

// 	if !decRegex.MatchString(r.Total) {
// 		return "", throwFormatError(r.Total), nil
// 	}
// 	suffix := r.Total[len(r.Total)-2:]

// 	// 50 points if total amount is round dollar amount (ends in .00)
// 	// 25 points if total amount is a multiple of .25
// 	switch suffix {
// 	case "00":
// 		suffixPoints += 75
// 	case ".25", ".50", ".75":
// 		suffixPoints += 50
// 	}

// 	// 5 points per 2 items on receipt
// 	itemsLengthPoints += (len(r.Items) / 2) * 5

// 	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// 	for _, item := range r.Items {
// 		if (len(strings.Trim(item.ShortDescription, " ")))%3 == 0 {
// 			if !decRegex.MatchString(item.Price) {
// 				fmt.Println(throwFormatError(item.ShortDescription))
// 				continue
// 			}
// 			multipliedPrice, _ := strconv.ParseFloat(item.Price, 64)

// 			itemDescriptionPoints += int(math.Ceil(multipliedPrice * .2))
// 		}
// 	}

// 	// 6 points if the day in the purchased date is odd
// 	var dateLenErr = len(r.PurchaseDate) != 10
// 	var day, dateErr = strconv.Atoi(r.PurchaseDate[9:])
// 	if dateErr != nil || dateLenErr == true {
// 		return "", throwFormatError(r.PurchaseDate), nil
// 	}

// 	if day % 2 != 0 {
// 		dayPoints += 6
// 	}

// 	// 10 points if time of purchase is between 2:00 and 4:00
// 	if len(r.PurchaseTime) == 5 {
// 		var hours, hoursErr = strconv.Atoi(r.PurchaseTime[0:2])
// 		var minutes, minutesErr = strconv.Atoi(r.PurchaseTime[3:5])
// 		if hoursErr != nil || minutesErr != nil {
// 			return "", throwFormatError(r.PurchaseTime), nil
// 		}

// 		var time = (hours * 100) + minutes

// 		if time > 1400 && time < 1600 {
// 			timePoints += 10
// 		}
// 	}

// 	points = retailerPoints + suffixPoints + itemsLengthPoints + itemDescriptionPoints + dayPoints + timePoints
// 	pointsSlice := []int{retailerPoints, suffixPoints, itemsLengthPoints, itemDescriptionPoints, dayPoints, timePoints}
// 	return strconv.Itoa(points), nil, pointsSlice
// }
