package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1", nil)
	request.Header.Set("APIkey", "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb")
	response, err := client.Do(request)

	data1, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data1))
	if err != nil {
		fmt.Print(err)
	}

}
