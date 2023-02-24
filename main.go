package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	ical "github.com/arran4/golang-ical"
)

func main() {
	b, err := os.ReadFile("sample.ics")
	if err != nil {
		log.Fatal(err)
	}
	cal, err := ical.ParseCalendar(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}
	// full file
	fmt.Println(cal.Serialize())

	for _, v := range cal.Events() {
		fmt.Println(v)

	}
}
