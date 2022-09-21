package main

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

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




type Doctor struct {
	UserName  string `json:"UserName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string`json:"LastName"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    
	db, err := sql.Open("mysql", "DB_URL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	var doctor Doctor
	
	if err := json.Unmarshal([]byte(request.Body), &doctor); err != nil {
        return events.APIGatewayProxyResponse{}, err
    }



    
	return events.APIGatewayProxyResponse{Body: register(db,doctor),StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

func register(db *sql.DB,doc Doctor) string{





	h := sha256.New()
	h.Write([]byte(doc.Password))
	bs := h.Sum(nil)
	doc.Password = hex.EncodeToString(bs)

	_, err := db.Exec("INSERT INTO Doctor (UserName,Email,Password,FirstName,LastName) values ('" + doc.UserName + "' ,'" + doc.Email + "', '" + doc.Password + "', '" + doc.FirstName + "', '" + doc.LastName + "')")

	if err != nil {
		return "false"
	}
	return "true"

}
