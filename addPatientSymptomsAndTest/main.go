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
	Bleeding int `json:"Bleeding"`
    Bruising int `json:"Bruising"`
    Pregnant int `json:"Pregnant"`
    Biopsy int `json:"Biopsy"`
    ChestPain int `json:"ChestPain"`
    ChestInjury int `json:"ChestInjury"`
    ShortnessOfBreath int `json:"ShortnessOfBreath"`
    KidneyDisease int `json:"KidneyDisease"`
    LiverDisease int `json:"LiverDisease"`
    Inflammation int `json:"Inflammation"`
    Urinate int `json:"Urinate"`
    Thirsty int `json:"Thirsty"`
    LoseWeight int `json:"LoseWeight"`
    Hungry int `json:"Hungry"`
    BlurryVision int `json:"BlurryVision"`
    Numb int `json:"Numb"`
    Tired int `json:"Tired"`
    SoresHealSlowly int `json:"SoresHealSlowly"`
    Infections int `json:"Infections"`
	CBC int `json:"CBC"`
    Ultrasound int `json:"Ultrasound"`
    ChestX_rays int `json:"ChestX_rays"`
    BloodChemistry int `json:"BloodChemistry"`
    Bloodglucose int `json:"Bloodglucose"`
	Date string `json:"Date"`

}


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    
	db, err := sql.Open("mysql", "DB_URL")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	var pSAT PatientSymptomsAndTest
	
	if err := json.Unmarshal([]byte(request.Body), &pSAT); err != nil {
        return events.APIGatewayProxyResponse{}, err
    }



    
	return events.APIGatewayProxyResponse{Body: addPatientSymptomsAndTest(db,pSAT),StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

func addPatientSymptomsAndTest(db *sql.DB,pSAT PatientSymptomsAndTest) string{





	_, err := db.Exec("INSERT INTO PatientSymptomsAndTest (patientID, DoctorUserName, bleeding, bruising, pregnant, biopsy, chestPain, chestInjury, shortnessOfBreath, KidneyDisease, LiverDisease, Inflammation, Urinate, thirsty, LoseWeight, hungry, blurryVision, numb, tired, soresHealSlowly, infections, CBC, ultrasound, ChestX_rays, bloodChemistry, Bloodglucose, date) VALUES ("+strconv.Itoa(pSAT.PatientID)+",'"+pSAT.DoctorUserName+"',"+strconv.Itoa(pSAT.Bleeding)+","+strconv.Itoa(pSAT.Bruising)+","+strconv.Itoa(pSAT.Pregnant)+","+strconv.Itoa(pSAT.Biopsy)+","+strconv.Itoa(pSAT.ChestPain)+","+strconv.Itoa(pSAT.ChestInjury)+","+strconv.Itoa(pSAT.ShortnessOfBreath)+","+strconv.Itoa(pSAT.KidneyDisease)+","+strconv.Itoa(pSAT.LiverDisease)+","+strconv.Itoa(pSAT.Inflammation)+","+strconv.Itoa(pSAT.Urinate)+","+strconv.Itoa(pSAT.Thirsty)+","+strconv.Itoa(pSAT.LoseWeight)+","+strconv.Itoa(pSAT.Hungry)+","+strconv.Itoa(pSAT.BlurryVision)+","+strconv.Itoa(pSAT.Numb)+","+strconv.Itoa(pSAT.Tired)+","+strconv.Itoa(pSAT.SoresHealSlowly)+","+strconv.Itoa(pSAT.Infections)+","+strconv.Itoa(pSAT.CBC)+","+strconv.Itoa(pSAT.Ultrasound)+","+strconv.Itoa(pSAT.ChestX_rays)+","+strconv.Itoa(pSAT.BloodChemistry)+","+strconv.Itoa(pSAT.Bloodglucose)+",'"+pSAT.Date+"');")

	if err != nil {
		return "false"
	}
	return "true"

}
