/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"github.com/morgulbrut/transportCli/webreq"

	"github.com/spf13/cobra"
)

// coordinatesCmd represents the coordinates command
var coordinatesCmd = &cobra.Command{
	Use:   "coordinates",
	Short: "Returns nearby stations trough coordinates.",
	Long:  `Coordinates in lat lon format, returns nearby stations`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintOut(webreq.Location("?x=" + args[0] + "&y=" + args[1]))
	},
}

func init() {
	locationCmd.AddCommand(coordinatesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coordinatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coordinatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
