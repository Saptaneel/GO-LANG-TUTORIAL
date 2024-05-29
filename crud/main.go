package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	Userid    int    `json:"Userid"`
	Id        int    `json:"Id"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

func getreq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	/*data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(data))*/

	var tod Todo

	err = json.NewDecoder(res.Body).Decode(&tod)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(tod)
}

func postreq() {
	tod := Todo{
		Userid:    23,
		Title:     "Virat",
		Completed: true,
	}

	jsondata, err := json.Marshal(tod)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonstring := string(jsondata)

	jsonreader := strings.NewReader(jsonstring)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURL, "application/json", jsonreader)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

	fmt.Println("Response status : ", res.Status)

}

func updatereq() {
	tod := Todo{
		Userid:    18,
		Title:     "Virat Kohli",
		Completed: true,
	}
	jsondata, err := json.Marshal(tod)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonstring := string(jsondata)

	jsonreader := strings.NewReader(jsonstring)
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, url, jsonreader)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	fmt.Println("Response status : ", res.Status)

}
func main() {
	/*getreq()*/
	/*postreq()*/
	updatereq()

}
