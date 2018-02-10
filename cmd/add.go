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
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/retr0h/logbook/client"
	"github.com/retr0h/logbook/entry"
	"github.com/spf13/cobra"
)

// AddCmd TODO
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new logbook entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.New(); err != nil {
			return err
		}

		callSign, name := getInput()
		ne := entry.NewEntry(callSign, name)

		err := client.Put([]byte(ne.CallSign), ne)
		if err != nil {
			return err
		}
		fmt.Println()
		color.Green("saved")

		client.Close()

		return nil
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}

func getInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Call Sign: ")
	scanner.Scan()
	callSign := scanner.Text()
	fmt.Print("Enter Name: ")
	scanner.Scan()
	name := scanner.Text()

	return callSign, name
}
