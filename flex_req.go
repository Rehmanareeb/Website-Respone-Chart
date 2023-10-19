package main

import (
	charts "flex/Charts"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Sending_Flex_Request() {

	Url := "https://flexstudent.nu.edu.pk/Login"

	request, err := http.NewRequest("GET", Url, nil)

	request.Header.Set("Content-Type", "application/json; charset= utf=8")

	Client := &http.Client{}

	response, err := Client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status for the sent request :", response.Status)
}

func main() {
	var each_time_taken []int64 = make([]int64, 10)

	for i := 0; i < 10; i++ {
		start_time := time.Now()
		Sending_Flex_Request()
		end_time := time.Now()

		Time_Taken := end_time.Sub(start_time)

		each_time_taken[i] = Time_Taken.Milliseconds()
	}
	fmt.Println("The total time taken by the request to respond was", each_time_taken)

	http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		charts.Generate_Bar_Chart(w, r, each_time_taken)
	})

	port := 3000
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
