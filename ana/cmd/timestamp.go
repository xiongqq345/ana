/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewTimestampCommand() *cobra.Command {
	tc := &cobra.Command{
		Use:   "time",
		Short: "Conversion between timestamp and time",
		Run:   timestampCommandFunc,
	}

	return tc
}

func timestampCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("time command needs only one argument"))
	}

	var t time.Time
	timeStr := args[0]
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		if t, err = parseTimeString(timeStr); err != nil {
			ExitWithError(ExitBadArgs, err)
		}
	} else {
		if timestamp < 1e11 {
			t = time.Unix(timestamp, 0)
		} else {
			t = time.Unix(timestamp/1e3, 0)
		}
	}
	fmt.Printf("time string  : %s\n", t.Format("2006-01-02 15:04:05"))
	fmt.Printf("timestamp(s) : %d\n", t.Unix())
	fmt.Printf("timestamp(ms): %d\n", t.UnixNano()/1e6)
	fmt.Printf("timestamp(μs): %d\n", t.UnixNano()/1e3)
	fmt.Printf("timestamp(ns): %d\n", t.UnixNano())
}

func parseTimeString(ts string) (tm time.Time, err error) {
	if ts == "now" {
		return time.Now(), nil
	}

	layouts := []string{
		"2006-01-02 15:04:05",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		tm, err = time.Parse(layout, ts)
		if err == nil {
			return tm, nil
		}
	}

	return time.Time{}, fmt.Errorf("unknown time layout")
}
