package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func calculatePoints(r *receipt) (string, error, []int) {
	retailerPoints, suffixPoints, itemsLengthPoints, itemDescriptionPoints, dayPoints, timePoints := 0, 0, 0, 0, 0, 0
	points := 0

	decRegex, _ := regexp.Compile(`^\d+\.?\d\d$`)
	alphaNumericRegex, _ := regexp.Compile(`[^a-zA-Z0-9 ]+`)

	
	formattedRetailName := formatString(alphaNumericRegex, r.Retailer)
	
	// calculate retail points
	retailerPoints = calcRetailPoints(formattedRetailName)

	if !decRegex.MatchString(r.Total) {
		return "", throwFormatError(r.Total), nil
	}
	suffix := r.Total[len(r.Total)-2:]

	// calculate suffix points
	suffixPoints = calcSuffixPoints(suffix)
	
	// calculate items[] length points	
	itemsLengthPoints = calcItemsLengthPoints(r.Items)

	// calculate item description points
	itemDescriptionPoints = calcItemDescriptionPoints(r.Items, *decRegex)
	
	// calculate date points
	dayPoints,_ = calcDayPoints(r.PurchaseDate)
	
	// calculate time points
	timePoints,_ = calcTimePoints(r.PurchaseTime)
	

	points = retailerPoints + suffixPoints + itemsLengthPoints + itemDescriptionPoints + dayPoints + timePoints
	pointsSlice := []int{retailerPoints, suffixPoints, itemsLengthPoints, itemDescriptionPoints, dayPoints, timePoints}
	return strconv.Itoa(points), nil, pointsSlice
}

// 1 point per alphanumeric character in retailer name
func calcRetailPoints(s string) int {
	return len(strings.ReplaceAll(s, " ", ""))
}

// 50 points if total amount is round dollar amount (ends in .00)
// 25 points if total amount is a multiple of .25
func calcSuffixPoints(s string) int {
	switch s {
	case "00":
		return 75
	case ".25", ".50", ".75":
		return 50
	}
	return 0
}

// 5 points per 2 items 
func calcItemsLengthPoints(i []item) int {
	return (len(i) / 2) * 5
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func calcItemDescriptionPoints(i []item, r regexp.Regexp) int {
	tempPoints := 0
	for _, item := range i {
		if (len(strings.Trim(item.ShortDescription, " ")))%3 == 0 {
			if !r.MatchString(item.Price) {
				fmt.Println(throwFormatError(item.ShortDescription))
				continue
			}
			multipliedPrice, _ := strconv.ParseFloat(item.Price, 64)

			tempPoints += int(math.Ceil(multipliedPrice * .2))
		}
	}
	return tempPoints
}

// 6 points if the day in the purchased date is odd
func calcDayPoints(s string) (int, error) {
	var dateLenErr = len(s) != 10
	var day, dateErr = strconv.Atoi(s[9:])
	if dateErr != nil || dateLenErr == true {
		return 0, throwFormatError(s)
	}

	if day % 2 != 0 {
		return 6, nil
	}
	return 0, nil
}

// 10 points if the time of purchase is after 2:00PM and before 4:00PM
func calcTimePoints(s string) (int, error) {
	if len(s) == 5 {
		var hours, hoursErr = strconv.Atoi(s[0:2])
		var minutes, minutesErr = strconv.Atoi(s[3:5])
		if hoursErr != nil || minutesErr != nil {
			return 0, throwFormatError(s)
		}

		var time = (hours * 100) + minutes

		if time > 1400 && time < 1600 {
			return 10, nil
		}
	}
	return 0, nil
}

func formatString(re *regexp.Regexp, s string) string {
	return re.ReplaceAllString(s, "")
}

func throwFormatError(s string) error {
	return errors.New("Incorrect format for " + s + "\n")
}