// Copyright (c) 2018 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/retr0h/logbook/client"
	"github.com/retr0h/logbook/entry"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	callSign string
)

// GetCmd TODO
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a particular logbook entry",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRequiredFlags(cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.New(); err != nil {
			return err
		}

		data, err := client.Get([]byte(callSign))
		if err != nil {
			return err
		}

		e, err := entry.Unmarshal(data)
		if err != nil {
			return err
		}

		timeLocalTZ, _ := time.Now().In(time.Local).Zone()
		timeFormatString := fmt.Sprintf("01/02/2006 15:04:05 %s", timeLocalTZ)
		createdAt := e.CreatedAt.Local().Format(timeFormatString)
		ID := strconv.Itoa(e.ID)
		tableData := [][]string{
			[]string{createdAt, e.CallSign, e.Name, ID},
		}

		fmt.Println("")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Call Sign", "Name", "ID"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(tableData)
		table.Render()
		fmt.Println("")

		client.Close()

		return nil
	},
}

func checkRequiredFlags(flags *pflag.FlagSet) error {
	missingFlagNames := []string{}

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			missingFlagNames = append(missingFlagNames, flag.Name)
		}
	})

	if len(missingFlagNames) > 0 {
		return fmt.Errorf("Required flag/flags: \"%s\" has/have not been set", strings.Join(missingFlagNames, "\", \""))
	}

	return nil
}

func init() {
	RootCmd.AddCommand(GetCmd)
	GetCmd.Flags().StringVarP(&callSign, "callsign", "c", "", "Users call sign")
	GetCmd.MarkFlagRequired("callsign")
}
