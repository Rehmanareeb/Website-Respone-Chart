package main

import (
	charts "flex/Charts"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Sending_Request_For_Website(Url string) {

	request, err := http.NewRequest("GET", Url, nil)

	request.Header.Set("Content-Type", "application/json; charset= utf=8")

	Client := &http.Client{}

	response, err := Client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println("Status for the sent request :", response.Status)
}

func main() {
	var each_time_taken []float64 = make([]float64, 10)
	var Total_Time_Taken float64

	Url := "https://" + charts.Get_The_Name_OfWebsite()

	for i := 0; i < 10; i++ {
		start_time := time.Now()
		Sending_Request_For_Website(Url)
		end_time := time.Now()

		Time_Taken := end_time.Sub(start_time)

		each_time_taken[i] = Time_Taken.Seconds()
		Total_Time_Taken += float64(each_time_taken[i])
	}
	fmt.Println("The total time taken by the request to respond was", Total_Time_Taken)

	http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		charts.Generate_Bar_Chart(w, r, each_time_taken)
	})

	port := 3000
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

	url := charts.Get_The_Name_OfWebsite()
	fmt.Print(url)
}
