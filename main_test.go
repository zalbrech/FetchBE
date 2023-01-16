package main

import (
	"github.com/gin-gonic/gin"
)

// 28 points
var postReceipt1 = receipt {
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items:        testItems1,
	Total:        "35.35",
}  

// 109 points
var postReceipt2 = receipt {
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items:        testItems2,
	Total:        "9.00",
}

//TODO : finish post test
func testPost(cntx *gin.Context) {
	if err := cntx.BindJSON(&postReceipt1); err != nil {
		panic(err)
	}
}