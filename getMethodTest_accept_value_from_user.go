package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func myGet1(url *string, key *string) {

	fmt.Println("Inside the function " + url + "  " + key)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("APIkey", key)
	fmt.Println("Request is : ", request)
	response, err1 := client.Do(request)

	if err != nil {
		fmt.Print(err)
	}

	if err1 != nil {
		fmt.Print(err)
	}
	fmt.Println("Responce is : ", response)
	data1, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		fmt.Println("Error in calling method : " + err2.Error())
	}
	fmt.Println(string(data1))

}

func main() {

	var url string
	var key string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter SAP API Service: ")
	url, _ = reader.ReadString('\n')
	fmt.Println(url)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter SAP APIKey : ")
	key, _ = reader1.ReadString('\n')
	fmt.Println(key)

	//var url1 string
	//	url1 = url

	//	var key1 string
	//key1 = key

	//	var url = "https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1"
	//	var key = "B6NePYsFuQRE9REfQiUCdtwcGANYOVsb"

	client := &http.Client{}
	//"https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1"
	//"B6NePYsFuQRE9REfQiUCdtwcGANYOVsb"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("APIkey", key)
	fmt.Println("Request is : ", request)
	response, err1 := client.Do(request)

	if err != nil {
		fmt.Print(err)
	}

	if err1 != nil {
		fmt.Print(err)
	}
	fmt.Println("Responce is : ", response)

	myGet1(&url, &key)
}
