package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)

	if err != nil {
		println(err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%s\n", body)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go fetch("http://www.golang.org", &wg)
	go fetch("http://www.google.com", &wg)

	wg.Wait()

}
