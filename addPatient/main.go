package main

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
    "strconv"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)




type Patient struct {
	ID  int `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age  int `json:"Age"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    
	db, err := sql.Open("mysql", "DB_URL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	var patient Patient
	
	if err := json.Unmarshal([]byte(request.Body), &patient); err != nil {
        return events.APIGatewayProxyResponse{}, err
    }



    
	return events.APIGatewayProxyResponse{Body: addpatient(db,patient),StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

func addpatient(db *sql.DB,patient Patient) string{





	_, err := db.Exec("INSERT INTO patient (ID,FirstName,LastName,Age) values ("+strconv.Itoa(patient.ID)+",'"+patient.FirstName+"','"+patient.LastName+"',"+strconv.Itoa(patient.Age)+")")

	if err != nil {
		return "false"
	}
	return "true"

}
