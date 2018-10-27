package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//responce, err := http.ServerContextKey("B6NePYsFuQRE9REfQiUCdtwcGANYOVsb")
	//http.Request.api.ServerContextKey:= "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb"
	//http.Header.Set("APIkey", "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb", "")
	//http.Header.Set.APIkey :=  "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb"
	//https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1
	responce, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Println("There is some error as %s \n", err)
	} else {
		data, _ := ioutil.ReadAll(responce.Body)
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"name": "NICE"}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("GET", "https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1", bytes.NewBuffer(jsonValue))
	request.Header.Set("APIkey", "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb")
	data1, _ := ioutil.ReadAll(request.Body)
	fmt.Println(string(data1))

}
