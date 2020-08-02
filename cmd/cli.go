package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	sbb_api "github.com/denysvitali/go-sbb-api/pkg"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var args struct {
	Date string `arg:"-d"`
	Time string `arg:"-t"`
	From string `arg:"positional, required"`
	To   string `arg:"positional, required"`
}

func main() {
	arg.MustParse(&args)

	client := sbb_api.New()

	if args.Date == "" {
		args.Date = time.Now().Format("2006-01-02")
	}

	if args.Time == "" {
		args.Time = time.Now().Format("15:04:05")
	}

	at_date, err := time.Parse("2006-01-02", args.Date)
	if err != nil {
		logrus.Fatal(err)
	}
	at_time, err := time.Parse("15:04:05", args.Time)
	if err != nil {
		logrus.Fatal(err)
	}

	year, month, day := at_date.Date()

	hour, minute, second := at_time.Hour(), at_time.Minute(), at_time.Second()

	at := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	connections, err := client.GetConnections(args.From, args.To, at)

	if err != nil {
		logrus.Fatal(err)
	}

	for _, conn := range connections.Connections {
		fmt.Printf("Connection %s\n", conn.TransportDesignation.TransportText)
		if len(conn.Via) != 0 {
			fmt.Printf("\tVia: %s\n", strings.Join(conn.Via, ","))
		}
		fmt.Printf("\tDeparture: %s (%s) - %s %s \n", conn.Departure, conn.DepartureTime, conn.DepartureTrackLabel, conn.DepartureTrack)
		fmt.Printf("\tArrival: %s (%s)\n", conn.Destination, conn.ArrivalTime)
		fmt.Printf("\tDuration: %s\n", conn.Duration)
		fmt.Printf("\tTransfers: %d\n", conn.Transfers)
		for _, section := range conn.Sections {
			fmt.Printf("\n")
			fmt.Printf("\t%s (%s) - %s %s\n", section.DepartureName, section.DepartureTime,
				section.DepartureTrackLabel, section.DepartureTrack)
			fmt.Printf("\t%s (%s) - %s %s\n", section.ArrivalName, section.ArrivalTime,
				section.ArrivalTrackLabel, section.ArrivalTrack)
			if section.Type == "TRANSPORT" {
				fmt.Printf("\t%s: %s (%s)\n",
					section.TransportDesignation.TransportName,
					section.TransportDesignation.TransportLabel,
					section.TransportDesignation.TransportText)
			} else if section.Type == "WALK" {
				fmt.Printf("\tWalk from %s to %s (%s)",
					section.DepartureName,
					section.ArrivalName,
					section.WalkDescription,
					)
			} else {
				fmt.Printf("\tSection Type: %s\n", section.Type)
			}
		}
		fmt.Printf("\n")
	}
}
