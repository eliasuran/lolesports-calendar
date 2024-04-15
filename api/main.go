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

	fmt.Println("Calendars:")
	for _, item := range calendarList.Items {
		fmt.Println(item.Id)
	}

}
