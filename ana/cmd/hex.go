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

	"github.com/spf13/cobra"
)

func NewHexCommand() *cobra.Command {
	lc := &cobra.Command{
		Use:   "hex",
		Short: "Prints the arguments length",
		Run:   hexCommandFunc,
	}

	return lc
}

func hexCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("hex command needs only one argument"))
	}

	s, _ := strconv.Unquote(`"` + args[0] + `"`)
	fmt.Println(s)
}
