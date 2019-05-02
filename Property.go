package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tenant-management/model"
)

//var ownerList = make(map[string]model.Owner)
var propertyList = make(map[string]Property)
//var tenantList = make(map[string]Tenant)

func main() {

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8080")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/updateProperty", updateProperty)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func updateProperty(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var property model.Property
		errM := json.Unmarshal(data, &property)

		// save data
		propertyList[property.propertyName] = property

		if errM != nil {
			http.Error(w, errM.Error(), 500)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "Property Updated %s!",  property.propertyName)
	}
}
