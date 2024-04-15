package main

import (
	"context"
	"fmt"
	"os"

	"github.com/eliasuran/lolesports-calendar-api/oauth"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	client := oauth.Authorize()

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Printf("Unable to retrieve Calendar client: %v\n", err)
		os.Exit(1)
	}

	calendarList, err := srv.CalendarList.List().Do()
	if err != nil {
		fmt.Printf("Could not get calendar list: %v\n", err)
		os.Exit(1)
	}

	event := &calendar.Event{
		Summary:     "Google I/O 2015",
		Location:    "800 Howard St., San Francisco, CA 94103",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2015-05-28T09:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2015-05-28T17:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: "lpage@example.com"},
			&calendar.EventAttendee{Email: "sbrin@example.com"},
		},
	}

	calendarId := calendarList.Items[0].Id
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		fmt.Printf("Couldnt create event: %v\n", err)
	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)
}
