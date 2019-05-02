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
fmt.Println("Starting the application...")
    response, err := http.Get(propertyList[property.propertyName])
    if err != nil {
        fmt.Printf("The value didn't get %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"ownerEmail": "sanniraj@gmail.com", "propertyName": "Sanni Raj"}
    jsonValue, _ := json.Marshal(jsonData)
    fmt.Println("Terminating the application...")
}
}
