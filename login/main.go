package main

import (
	"errors"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"crypto/sha256"
	"encoding/hex"
	

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





func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := sql.Open("mysql", "DB_URL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	usernmae:= request.QueryStringParameters["username"]
	password:= request.QueryStringParameters["password"]
	

	return events.APIGatewayProxyResponse{
		Body:       login(db,usernmae,password),
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(handler)
}

func login(db *sql.DB,UserName string, Password string) string {

	h := sha256.New()

	h.Write([]byte(Password))
	bs := h.Sum(nil)
	Password = hex.EncodeToString(bs)

	var usernmae string
	var password string

	err := db.QueryRow("SELECT UserName,Password from Doctor where UserName='"+UserName+"' and Password='"+Password+"' ;").Scan(&usernmae, &password)
	if err != nil {
		return "false"
	}

	return "true"

}
