package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	var client http.Client
	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Erro")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro2")
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		//log.Info(bodyString)
	}
}
