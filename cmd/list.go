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

	"github.com/olekukonko/tablewriter"
	"github.com/retr0h/logbook/client"
	"github.com/retr0h/logbook/entry"
	"github.com/spf13/cobra"
)

// ListCmd TODO
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all logbook entries",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRequiredFlags(cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.New(); err != nil {
			return err
		}

		var data [][]string

		for _, bp := range client.List() {
			e, err := entry.Unmarshal(bp.Value)
			if err != nil {
				return err
			}

			ID := strconv.Itoa(e.ID)
			data = append(data, []string{ID, e.CallSign})
		}

		fmt.Println("")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Call Sign"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data)
		table.Render()
		fmt.Println("")

		client.Close()

		return nil
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
