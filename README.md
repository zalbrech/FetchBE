# FetchBE
This repository defines a basic API with a few endpoints. 

# To Run Locally


## Check Go Verssion
Use ```go version``` to check that Go is installed on your machine

#### Note that you need v1.16+ to run this application

If it is not installed, Go can be downloaded and installed here:

    https://go.dev/doc/install

## Clone Repository 

Clone the repository with the following link:

    https://github.com/zalbrech/FetchBE.git
    
## Run application    
    
Navigate to the local FetchBE directory, then run ```go run .``` 

# API

The API is described below

### Request

`GET /receipts`

    curl --location --request GET "http://localhost:8080/receipts" 
### Resoponse

Array of JSON objects
<br/><br/>

### Request

`GET /receipts/{id}/points`

    curl --location --request GET "http://localhost:8080/receipts/{id}/points"
    
### Response

Returns the points associated with the {id} param
<br/><br/>


### Request

`POST /receipts`

    curl --location --request POST "localhost:8080/receipts" --header "Content-Type: application/json" --data-raw "{data}"   
    
### Response

Returns a generated UUID that can be used to find the points for the receipt associated with that id
<br/><br/>

## Examples

### POST

#### Double Quote (Windows Powershell)

    curl -L -X POST "localhost:8080/receipts" -H "Content-Type: application/json" --data-raw "{\"retailer\":\"Woodman's\",\"purchaseDate\":\"2022-03-20\",\"purchaseTime\":\"14:33\",\"items\":[{\"shortDescription\":\"Gatorade\",\"price\":\"2.25\"},{\"shortDescription\":\"Gatorade\",\"price\":\"2.25\"},{\"shortDescription\":\"Gatorade\",\"price\":\"2.25\"},{\"shortDescription\":\"Gatorade\",\"price\":\"2.25\"}],\"total\":\"9.00\"}"
    
#### Single Quotes

    curl -L -X POST 'localhost:8080/receipts' -H 'Content-Type: application/json' --data-raw '{"retailer":"Woodman'\''s","purchaseDate":"2022-03-20","purchaseTime":"14:33","items":[{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"}],"total":"9.00"}'


### Response

   "id: 0c4787b2-7361-4fc0-9590-5fcff262bd23"
   
#### Note this is sample data, the generated id will be different
<br/><br/>


### GET Points

    curl --location --request GET "http://localhost:8080/receipts/0c4787b2-7361-4fc0-9590-5fcff262bd23/points"

### Response

    "109"