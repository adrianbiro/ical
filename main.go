package main

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

func main() {
	/*
		b, err := os.ReadFile("sample.ics")
		if err != nil {
			log.Fatal(err)
		}
		cal, err := ics.ParseCalendar(strings.NewReader(string(b)))
		if err != nil {
			log.Fatal(err)
		}
		// full file
		fmt.Println(cal1.Serialize())
	*/
	//
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf("id@domain"))
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(time.Now())
	event.SetModifiedAt(time.Now())
	event.SetStartAt(time.Now())
	event.SetEndAt(time.Now())
	event.SetSummary("Summary bum bum")
	event.SetLocation("Address Kapurkova 5")
	event.SetDescription("Description foo bar hola hej asdasdasd987987987")
	event.SetURL("https://google.com")
	event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
	event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
	event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
	fmt.Println(cal.Serialize())
}
