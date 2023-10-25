package charts

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Fetching_Data_Using_The_Time(time_taken []float64) []opts.BarData {

	var data []opts.BarData

	for i := 0; i < len(time_taken); i++ {

		data = append(data, opts.BarData{Value: time_taken[i]})

	}
	return data
}

func Generate_Bar_Chart(w http.ResponseWriter, _ *http.Request, time_taken []float64) {

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Time Taken By The Request To Respond",
		Subtitle: "In Seconds",
	}))

	bar.SetXAxis([]string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th", "10th"}).
		AddSeries("Request/Response ratio", Fetching_Data_Using_The_Time(time_taken))

	bar.Render(w)

}
