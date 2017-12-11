package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	DBUser, DBPass, DBName := GetSettings()
	db, err := sql.Open("mysql", DBUser+":"+DBPass+DBName)
	checkErr(err)
	defer db.Close() //Defer functions is called after return statements/when its finished

	stmtOut, err := db.Prepare("SELECT * FROM team WHERE id = 2")
	defer stmtOut.Close()

    
	fmt.Print(DBUser, DBPass, DBName)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Gets the db user information from external json file.
func GetSettings() (string, string, string) {
	var settings = new(Settings)
	raw, err := ioutil.ReadFile("dbSettings.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(raw, &settings)
	if err != nil {
		return "", "", ""
	}
	return settings.DBUser, settings.DBPass, settings.DBName
}
