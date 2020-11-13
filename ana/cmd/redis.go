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

func NewRedisCommand() *cobra.Command {
	tc := &cobra.Command{
		Use:   "redis",
		Short: "Redis help",
		Run:   redisCommandFunc,
	}

	return tc
}

func redisCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("time command needs only one argument"))
	}

	var tm time.Time
	timeArg := args[0]
	timestamp, err := strconv.ParseInt(timeArg, 10, 64)
	if err != nil {
		tm, err = checkTimeString(timeArg)
		if err != nil {
			ExitWithError(ExitBadArgs, err)
		}
	} else {
		if timestamp < 1e11 {
			tm = time.Unix(timestamp, 0)
		} else {
			tm = time.Unix(0, timestamp*1e6)
		}
	}
	fmt.Printf("time string  : %s\n", tm.Format("2006-01-02 15:04:05"))
	fmt.Printf("timestamp(s) : %d\n", tm.Unix())
	fmt.Printf("timestamp(ms): %d\n", tm.UnixNano()/1e6)
	fmt.Printf("timestamp(μs): %d\n", tm.UnixNano()/1e3)
	fmt.Printf("timestamp(ns): %d\n", tm.UnixNano())
}
