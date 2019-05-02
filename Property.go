package main

import (
	"fmt"
	"log"
	"net/http"
)

type Property struct {
	ownerEmail    string
	propertyName  string
	availableFlat int
	occupiedFlat  int
}

var propertyList = make(map[string]Property)

func main() {

	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8080")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/updateproperty", updateproperty)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func updateproperty(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var ownerEmail string
		var propertyName string
    var availableFlat int
    var occupiedFlat int 

		fmt.Println("Enter owner email:")
		_, _ = fmt.Scanf("%s", ownerEmail)

		fmt.Println("Enter owner propertyName:")
		_, _ = fmt.Scanf("%s", propertyName)

		property := Property{
			ownerEmail: ownerEmail,
			propertyName: propertyName,
      availableFlat: availableFlat,
      occupiedFlat: occupiedFlat,
		}

		propertyList[ownerEmail] = property
		fmt.Printf("Welcome %v", propertyName)
	}
}
