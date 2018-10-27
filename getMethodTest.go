package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func myGet(url string, key string) {

	fmt.Println("Inside the function")

	client := &http.Client{}
	//"https://sandbox.api.sap.com/successfactors/odata/v2/Background_Community?%24top=1"
	//"B6NePYsFuQRE9REfQiUCdtwcGANYOVsb"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("APIkey", key)
	response, err1 := client.Do(request)

	if err != nil {
		fmt.Print(err)
	}

	if err1 != nil {
		fmt.Print(err)
	}

	data1, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data1))

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter SAP API Service: ")
	url, _ := reader.ReadString('\n')
	fmt.Println(url)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter SAP APIKey : ")
	key, _ := reader1.ReadString('\n')
	fmt.Println(key)

	myGet(string(url), string(key))
}
