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
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

var historyNum int64

func NewHistoryCommand() *cobra.Command {
	hc := &cobra.Command{
		Use:   "history",
		Short: "Prints the command history",
		Run:   historyCommandFunc,
	}

	hc.Flags().Int64VarP(&historyNum, "num", "n", 20, "print last [n] commands")
	return hc
}

func historyCommandFunc(cmd *cobra.Command, args []string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:30637",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var buffer bytes.Buffer
	buffer.WriteString(cmd.CommandPath())
	if historyNum != 20 {
		buffer.WriteString(" --num ")
		buffer.WriteString(strconv.Itoa(int(historyNum)))
	}

	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: buffer.String(),
	}
	rdb.ZAdd(commandHistoryKey, member)

	val,err := rdb.ZRevRange(commandHistoryKey,0,historyNum).Result()
	if err != nil {
		fmt.Println(fmt.Errorf("redis error: %s",err))
	}
	for _,command := range val {
		fmt.Println(command)
	}
}
