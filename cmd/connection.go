/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/morgulbrut/transportCli/webreq"
	"github.com/morgulbrut/transportCli/webreq/parsejson"
	"github.com/spf13/cobra"
)

// connectionCmd represents the connection command
var connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Returns connections between two station",
	Long: `
    _____                                          _____     _______________ ________
    __  /________________ ________________ __________  /_    __  ____/___  / ____  _/
    _  __/__  ___/__  __ \__  ___/___  __ \__  ___/_  __/    _  /     __  /   __  /  
    / /_  _  /    _  / / /_(__  ) __  /_/ /_  /    / /_      / /___   _  /_____/ /   
    \__/  /_/     /_/ /_/ /____/  _  .___/ /_/     \__/      \____/   /_____//___/   
                                  /_/                                          
Returns connections between two station, needs two stations.

Stationnames longer than one word must be written in quotation marks. 

	Example: 	transportCli station "Bad Ragaz" Zürich
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var params strings.Builder
		if len(args) == 2 {
			params.WriteString("?from=" + args[0] + "&to=" + args[1])
		} else if len(args) == 1 {
			params.WriteString("?station=" + args[0])
		} else {
			cmd.Help()
		}

		lim, _ := cmd.Flags().GetString("limit")
		if lim != "" {
			params.WriteString("&limit=" + lim)
		} else { // default
			params.WriteString("&limit=1")
		}

		time, _ := cmd.Flags().GetString("time")
		if time != "" {
			params.WriteString("&time=" + time)
		}

		if len(args) == 2 {
			PrintConnection(webreq.Connection(params.String()))
		} else if len(args) == 1 {
			PrintStation(webreq.Station(params.String()))
		}
	},
}

func PrintConnection(resp parsejson.RespConnection) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if runtime.GOOS == "windows" {
		t.SetStyle(table.StyleDouble)
	} else {
		t.SetStyle(table.StyleColoredDark)
	}
	//t.AppendHeader(table.Row{"Departure", "Time", "Platform", "Arrival", "Time", "Platform", "Duration", "Changes"})

	t.AppendHeader(table.Row{"Departure", "Time", "P.", "Arrival", "Time", "P."})

	for _, ele := range resp.Connections {
		t.AppendSeparator()
		tfs := "2006-01-02T15:04:05-0700"
		for _, sec := range ele.Sections {
			td, _ := time.Parse(tfs, sec.Departure.Departure)
			ta, _ := time.Parse(tfs, sec.Arrival.Arrival)
			tds := fmt.Sprintf("%02d:%02d", td.Hour(), td.Minute())
			tas := fmt.Sprintf("%02d:%02d", ta.Hour(), ta.Minute())
			//t.AppendRow(table.Row{tms, ele.To, ele.PassList[0].Platform, ele.Category, ele.Number})
			//t.AppendRow(table.Row{ele.From.Station.Name, tds, ele.From.Platform, ele.To.Station.Name, tas, ele.To.Platform, durs, ele.Sections})
			t.AppendRow(table.Row{sec.Departure.Station.Name, tds, sec.Departure.Platform, sec.Arrival.Station.Name, tas, sec.Arrival.Platform})
		}
		t.AppendRow(table.Row{"Transfers", ele.Transfers, "", "Duration", ele.Duration})
	}
	t.Render()
}

func init() {
	rootCmd.AddCommand(connectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	connectionCmd.Flags().StringP("limit", "l", "", "Number of departing connections to return.")
	connectionCmd.Flags().StringP("time", "t", "", "Time of the earliest connection.")

}
