package main

import (
	// "regexp"
	"fmt"
	"strconv"
	"testing"
)

var testItems1 = []item{
	{
		ShortDescription: "Mountain Dew 12PK",
		Price:            "6.49",
	},
	{
		ShortDescription: "Emils Cheese Pizza",
		Price:            "12.25",
	},
	{
		ShortDescription: "Knorr Creamy Chicken",
		Price:            "1.26",
	},
	{
		ShortDescription: "Doritos Nacho Cheese",
		Price:            "3.35",
	},
	{
		ShortDescription: "     Klarbrunn 12-PK 12 FL OZ",
		Price:            "12.00",
	},
}

var testItems2 = []item {
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
	{
		ShortDescription: "Gatorade",
		Price: "2.25",
	},
}

// 28 points
var testReceipt1 = receipt {
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items:        testItems1,
	Total:        "35.35",
}  

// 109 points
var testReceipt2 = receipt {
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items:        testItems2,
	Total:        "9.00",
}


func TestTotalPoints(t *testing.T) {
	var tests = []struct {
		r receipt
		i []item
		want int
	}{
		{testReceipt1, testItems1, 28},
		{testReceipt2, testItems2, 109},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("receipt%d, %d", i+1, test.want)
		t.Run(testName, func(t *testing.T) {
			res,_ := calculatePoints(&test.r)
			numRes,_ := strconv.Atoi(res)
			if(numRes != test.want) {
				t.Errorf("got %d, want %d", numRes, test.want)
			}
		})
	}

    points,err := calculatePoints(&testReceipt1)
	want := 28
	numPoints,_ := strconv.Atoi(points)
	if(err != nil || numPoints != want) {
		t.Errorf("got %d, want %d", numPoints, want)
	}
}

func TestIndividualPoints(t *testing.T) {
	var tests = []struct {
		r receipt
		i []item
		retailerWant, suffixWant, itemsLengtWant, itemDescriptionWant, dayWant, timeWant int
	}{
		{testReceipt1, testItems1, 6, 0, 10, 6, 6, 0},
		{testReceipt2, testItems2, 14, 75, 10, 0, 0, 10},
	}

	for
}