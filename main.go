package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/bradhe/stopwatch"
)

func main() {

	channel := make(chan int)
	go worker(1, channel)
	go worker(2, channel)

	watch := stopwatch.Start()
	for i := 0; i < 10000; i++ {
		channel <- i
	}
	watch.Stop()
	fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())

	// for i := 0; i < 119; i++ {
	// 	go callApi(i)
	// }

	//waiting for threads
	//time.Sleep(20 * time.Second)

	/*
		1000 | 1 channel | 4.792 seconds
		1000 | 2 channel | 4.157 seconds
		1000 | 3 channel | 3.906 seconds
		10000 | 10 channel | 42 seconds
		10000 | 1 channel | 59 seconds
	*/

}

func worker(workerId int, number chan int) {
	for x := range number {
		callApi(workerId, x)
	}
}

func callApi(worker int, time int) {
	apiUrl := "https://localhost:7096/api/job"

	userData := []byte(`{ "keyword": "` + getProgrammingLanguageRandomly() + `", "country": "` + getCountryRandomly() + `" }`)

	// create new http request
	request, error := http.NewRequest("POST", apiUrl, bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// send the request
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(`Status: ` + response.Status + ` Worker: ` + fmt.Sprint(worker) + ` Time: ` + fmt.Sprint(time))

	// clean up memory after execution
	defer response.Body.Close()
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

func getProgrammingLanguageRandomly() string {
	programmingLanguages := []string{
		"Java",
		"Python",
		"JavaScript",
		"C++",
		"Ruby",
		"Go",
		"Swift",
		"Kotlin",
		"Rust",
		"PHP",
	}

	rand.Seed(time.Now().UnixNano())
	return programmingLanguages[rand.Intn(len(programmingLanguages))]
}

func getCountryRandomly() string {
	countries := []string{
		"India",
		"USA",
		"UK",
		"Canada",
		"France",
		"Germany",
		"Japan",
		"China",
		"Russia",
		"South Korea",
		"Brasil",
	}

	rand.Seed(time.Now().UnixNano())
	return countries[rand.Intn(len(countries))]
}
