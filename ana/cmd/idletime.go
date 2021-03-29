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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

var (
	itHost       string
	itPass       string
	itPattern    string
	itCreateTime string
	itDeviation  int64
)

type IdleTimeParse struct {
	CreateTime time.Time `json:"createTime,omitempty"`
}

func NewIdleTimeCommand() *cobra.Command {
	itc := &cobra.Command{
		Use:   "idletime",
		Short: "Idletime analysis for redis key prefix",
		Run:   IdleTimeCommandFunc,
	}

	itc.Flags().StringVarP(&itHost, "host", "h", "localhost:6379", "redis host")
	itc.Flags().StringVarP(&itPass, "auth", "a", "", "redis pass")
	itc.Flags().StringVarP(&itPattern, "pattern", "p", "", "key pattern")
	itc.Flags().StringVarP(&itCreateTime, "ct", "", "", "create time")
	itc.Flags().Int64VarP(&itDeviation, "deviation", "d", 14400, "create time deviation")
	return itc
}

func IdleTimeCommandFunc(cmd *cobra.Command, args []string) {
	var (
		ct  time.Time
		err error
	)
	if itPattern == "" {
		ExitWithError(ExitBadArgs, fmt.Errorf("key pattern must be set"))
	}
	if itCreateTime != "" {
		if ct, err = parseTimeString(itCreateTime); err != nil {
			ExitWithError(ExitBadArgs, err)
		}
	}
	rds := redis.NewClient(&redis.Options{
		Addr:     itHost,
		Password: itPass,
		DB:       0, // use default DB
	})
	keys, err := ScanKeys(rds, itPattern, nil)
	if err != nil {
		ExitWithError(ExitRedisError, err)
	}
	var itParse IdleTimeParse
	f, err := os.Create("")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	for _, key := range keys {
		val, err := rds.Get(key).Bytes()
		if err != nil {
			ExitWithError(ExitRedisError, err)
		}
		if !ct.Equal(time.Time{}) {
			if err := json.Unmarshal(val, &itParse); err != nil {
				fmt.Println(err)
			} else {
				write(f, " "+itParse.CreateTime.String())
			}
		}
		idleTime, err := rds.ObjectIdleTime(key).Result()
		if err != nil {
			ExitWithError(ExitRedisError, err)
		}
		write(f, idleTime.String()+"\n")
	}
}

func write(f *os.File, info string) {
	if _, err := f.Write([]byte(info)); err != nil {
		fmt.Println(err)
	}
}
