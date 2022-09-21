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




type PatientSymptomsAndTest struct {
	PatientID int `json:"PatientID"`
	DoctorUserName string `json:"DoctorUserName"`
	Bleeding int `json:"bleeding"`
    Bruising int `json:"bruising"`
    Pregnant int `json:"pregnant"`
    Biopsy int `json:"biopsy"`
    ChestPain int `json:"chestPain"`
    ChestInjury int `json:"chestInjury"`
    ShortnessOfBreath int `json:"shortnessOfBreath"`
    KidneyDisease int `json:"KidneyDisease"`
    LiverDisease int `json:"LiverDisease"`
    Inflammation int `json:"Inflammation"`
    Urinate int `json:"Urinate"`
    Thirsty int `json:"thirsty"`
    LoseWeight int `json:"LoseWeight"`
    Hungry int `json:"hungry"`
    BlurryVision int `json:"blurryVision"`
    Numb int `json:"numb"`
    Tired int `json:"tired"`
    SoresHealSlowly int `json:"soresHealSlowly"`
    Infections int `json:"infections"`
	CBC int `json:"CBC"`
    Ultrasound int `json:"ultrasound"`
    ChestX_rays int `json:"chestX_rays"`
    BloodChemistry int `json:"bloodChemistry"`
    Bloodglucose int `json:"bloodglucose"`
	Date string `json:"date"`

}


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    id,_:= strconv.Atoi(request.QueryStringParameters["ID"])
	
	db, err := sql.Open("mysql", "DB_URL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

    var history []PatientSymptomsAndTest=getPatientHistory(db,id)
    ResponseBody, _ :=json.Marshal(history)
	

    
	return events.APIGatewayProxyResponse{Body: string(ResponseBody),StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

func getPatientHistory(db *sql.DB,id int)[]PatientSymptomsAndTest {


	var history []PatientSymptomsAndTest

    
	results, err := db.Query("SELECT * FROM PatientSymptomsAndTest WHERE PatientID="+strconv.Itoa(id)+";")
    if err !=nil {
        return history
    }
    for results.Next() {
        var data PatientSymptomsAndTest
        err = results.Scan(&data.PatientID,&data.DoctorUserName,&data.Bleeding,&data.Bruising,&data.Pregnant,&data.Biopsy,&data.ChestPain,&data.ChestInjury,&data.ShortnessOfBreath,&data.KidneyDisease,&data.LiverDisease,&data.Inflammation,&data.Urinate,&data.Thirsty,&data.LoseWeight,&data.Hungry,&data.BlurryVision,&data.Numb,&data.Tired,&data.SoresHealSlowly,&data.Infections,&data.CBC,&data.Ultrasound,&data.ChestX_rays,&data.BloodChemistry,&data.Bloodglucose,&data.Date)

		if err !=nil {
            return(history)
        }
		history=append(history,data);
    }
	return history

}
