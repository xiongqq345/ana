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
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

func NewLenCommand() *cobra.Command {
	lc := &cobra.Command{
		Use:   "len",
		Short: "Prints the arguments length",
		Run:   lenCommandFunc,
	}

	return lc
}

func lenCommandFunc(cmd *cobra.Command, args []string) {
	str := strings.Join(args,"")
	length := len(str)
	count := utf8.RuneCountInString(str)
	fmt.Printf("length of string: %d\n",length)
	fmt.Printf("length of char:   %d\n",count)
}
