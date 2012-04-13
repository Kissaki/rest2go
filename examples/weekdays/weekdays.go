/*
	Example REST server

	When run can be opened in a webbrowser:
	http://localhost:3000/wd/ for index, showing all entries as logic of the Index function specifies
	http://localhost:3000/wd/We get full string for index We (other indices from map also possible)
*/
package main

import (
	"fmt"
	rest "github.com/Kissaki/rest2go"
	"log"
	"net/http"
)

var wdmap = map[string]string{
	"Mo": "Monday",
	"Tu": "Tuesday",
	"We": "Wednesday",
	"Th": "Thursday",
	"Fr": "Friday",
	"Sa": "Saturday",
	"Su": "Sunday",
}

type Weekdays struct {
	wd map[string]string
}

func (wd *Weekdays) Index(resp http.ResponseWriter) {
	for i, v := range wd.wd {
		fmt.Fprintf(resp, "%s: %s<br/>\n", i, v)
	}
}
func (wd *Weekdays) Find(resp http.ResponseWriter, id string) {
	if full, ok := wd.wd[id]; ok {
		fmt.Fprintf(resp, full)
	}
}

func main() {
	log.Println("Starting Server")
	address := "127.0.0.1:3000"

	var wd = new(Weekdays)
	wd.wd = wdmap
	rest.Resource("/wd/", wd)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalln(err)
	}
}
