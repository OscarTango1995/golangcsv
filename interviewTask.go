package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strconv"
)
const (
	username = "root"
	password = "idm444civic"
	hostname = "127.0.0.1:3306"
	dbname   = "interview_task"
	tablename= "persons"
)

func main() {
	//Opening and creating DB
	db := createAndOpenDB()

	// Open the file
	csvfile, err := os.Open("persons.csv")
	ErrorCheck(err,"CSV File Opened Successfully")
	// Reading the file
	readFile := csv.NewReader(csvfile)

	//SQL Query
	stmt, err := db.Prepare("INSERT INTO persons VALUES (?,?, ?, ?,?, ?, ?,?, ?, ?,?, ?, ?,?, ?, ?,?, ?, ?,?, ?, ?)")
	ErrorCheck(err,"No Errors in SQL Statement")

	i:=0
	success:=0
	// Iterating through the records
	for {
		// Reading each record from readFile
		record, err := readFile.Read()
		if err == io.EOF {
			//stop the loop when EOF reached
			break
		}
		if err != nil {
			//notify if any error comes during the process
			ErrorCheck(err,"Unexpected error occurred!")
			break
		}
		if i!=0 { //skipping the header
			age,_ :=strconv.Atoi(record[5])
			emp,_ := strconv.ParseBool(record[17])
			_,err =stmt.Exec(0,record[0], record[1],record[2], record[3],record[4], age,record[6], record[7],record[8], record[9],record[10], record[11],record[12], record[13],record[14], record[15],record[16], emp,record[18], record[19],record[20])
			success++
		}else {
			i++
		}
	}
	if success>0 {
		fmt.Printf("Success! Data Inserted Into '%s' Table \n",tablename)
		defer db.Close()
		ReadInsertedRecords()

	}else {
		fmt.Println("Oops!Something Weird Happened")
		defer db.Close()
	}
}

func createAndOpenDB() *sql.DB {

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+hostname+")/")
	ErrorCheck(err,"DB Server Connection Successful")
	PingDB(db)

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+dbname)
	ErrorCheck(err,dbname+" Database Created Successfully!")

	_,err = db.Exec("USE "+dbname)
	ErrorCheck(err,dbname+" Database Connection Established")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS "+tablename+" (" +
		"id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, first TEXT , last TEXT ," +
		" ssn TEXT , mother_tongue TEXT , race TEXT , age INT ," +
		" blood_group TEXT , gender TEXT , birthday TEXT , cc_number TEXT ," +
		" phone TEXT , email TEXT , street TEXT , state TEXT , city TEXT ," +
		" zip INT , graduated_from TEXT , employment_status BOOL , company TEXT ," +
		" designation TEXT , yearly_revenue TEXT)")
	ErrorCheck(err,tablename+ " Created Successfully!")

	_,err = db.Exec("USE "+dbname)
	ErrorCheck(err,dbname+" Connection Established")
	return db
}
func ReadInsertedRecords()  {
	db:=createAndOpenDB()
	rows, err := db.Query("SELECT * FROM "+tablename )
	ErrorCheck(err,"Nothing wrong with Query")
	defer rows.Close()
	fmt.Println("-------------------------------------------------------------------------------------------------------")
	for rows.Next() {
		var id, age, zip int
		var EmploymentStatus bool
		var first, last, ssn, MotherTongue, race, BloodGroup, gender,
		birthday, CcNumber, phone, email, street, state, city, GraduatedFrom,
		company, designation, YearlyRevenue string
		err = rows.Scan(&id, &first, &last, &ssn, &MotherTongue, &race,
			&age, &BloodGroup, &gender, &birthday, &CcNumber, &phone, &email,
			&street, &state, &city, &zip, &GraduatedFrom, &EmploymentStatus,
			&company, &designation, &YearlyRevenue)
		ErrorCheck(err,"")

		fmt.Printf("id: %d First Name: %s  Last Name: %s  SSN: %s  MotherTongue: %s  Race: %s  " +
			"Age: %d  BloodGroup: %s  Gender: %s  Birthday: %s  CcNumber: %s  Phone: %s  Email: %s  " +
			"Street: %s  State: %s  City: %s  Zip: %d  GraduatedFrom: %s  EmploymentStatus: %t  " +
			"Company: %s  Designation: %s  YearlyRevenue: %s \n\n",id, first, last, ssn, MotherTongue, race, age, BloodGroup, gender, birthday, CcNumber, phone, email, street, state, city, zip, GraduatedFrom, EmploymentStatus, company, designation, YearlyRevenue)
	}
	fmt.Println("*******************************************************************************************************")

	ErrorCheck(err,"Records Ended")
	defer db.Close()

}

func ErrorCheck(err error,message string) {
	if err != nil {
		panic(err.Error())
	}else {
		fmt.Println(message)
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err,"DB Server Ping Successful!")
}
