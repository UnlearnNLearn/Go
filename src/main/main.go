package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dnlo/struct2csv"
)

// http endpoint and file write operation,
// will accept one argument from command line and process that many records
// it 3 step process and we use go mod in this
// step 1 one http endpoint
// step 2 process the data
// step 3 save it to file
// endpoint to consume - https://random-data-api.com/api/users/random_user
//

// package main

type Person struct {
	Id                      uint   `json:"id"`
	Uid                     string `json:"uid"`
	Password                string `json:"password"`
	First_name              string `json:"first_name"`
	Last_name               string `json:"last_name"`
	Username                string `json:"username"`
	Email                   string `json:"email"`
	Avatar                  string `json:"avatar"`
	Gender                  string `json:"gender"`
	Phone_number            string `json:"phone_number"`
	Social_insurance_number string `json:"social_insurance_number"`
	Date_of_birth           string `json:"date_of_birth"`
	Employeeinfo            struct {
		Title     string `json:"title"`
		Key_skill string `json:"key_skill"`
	} `json:"employment"`
	Address struct {
		City           string `json:"city"`
		Street_name    string `json:"street_name"`
		Street_address string `json:"street_address"`
		Zip_code       string `json:"zip_code"`
		State          string `json:"state"`
		Country        string `json:"country"`
		Coordinates    struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	} `json:"address"`
	Credit_card struct {
		cc_number string
	} `json:"credit_card"`
	Subscription struct {
		Plan           string `json:"plan"`
		Status         string `json:"status"`
		Payment_method string `json:"payment_method"`
		Term           string `json:"term"`
	} `json:"subscription"`
}

func main() {
	start := time.Now()
	// read command line argument
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// open file
	var stats os.FileInfo
	var file1 *os.File
	file1, err = os.OpenFile("output.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer file1.Close()

	// get data from endpoint that many times
	var person Person
	csvWriter := csv.NewWriter(file1)
	var total int = 0
	for i := 0; i < n; i++ {
		resp, err := http.Get("https://random-data-api.com/api/users/random_user")
		if err != nil {
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		json.Unmarshal(body, &person)
		// filter data on the basis of male or female
		if person.Gender == "Male" || person.Gender == "Female" {
			enc := struct2csv.New()
			var rows [][]string
			headers, err := enc.GetColNames(person)
			if err != nil {
				return
			}
			row, err := enc.GetRow(person)
			if err != nil {
				return
			}
			stats, err = os.Stat("output.csv")
			if err != nil {
				return
			}
			if stats.Size() == 0 && total == 0 {
				rows = append(rows, headers)
			}
			rows = append(rows, row)

			for i := 0; i < len(rows); i++ {
				csvWriter.Write(rows[i])
			}
			total += 1
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println("processed ", total, " records")
	fmt.Println("time elapsed till now ", timeElapsed)

}
