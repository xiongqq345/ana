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

	"github.com/spf13/cobra"
)

var historyNum int16

func NewHistoryCommand() *cobra.Command {
	lc := &cobra.Command{
		Use:   "history",
		Short: "Prints the command history",
		Run:   historyCommandFunc,
	}

	lc.Flags().Int16VarP(&historyNum, "num", "n", 20, "print last [n] commands")
	return lc
}

func historyCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println(historyNum)
}
