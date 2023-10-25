package charts

import (
	"fmt"
	"strings"
)

func Get_The_Name_OfWebsite() string {
	var URL string

	fmt.Println("Enter the name of the Website on which you want to perform the response test: ")
	fmt.Scan(&URL)
	// Remove the newline character from the URL
	//URL = strings.TrimSuffix(URL, "\n")

	// Check if the URL has .com in it
	if !strings.HasSuffix(URL, ".com") {
		// Append .com to the URL
		URL = URL + ".com"

		return URL

	}
	return URL
}
