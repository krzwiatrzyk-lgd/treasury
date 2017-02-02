// Copyright © 2016 AirHelp Inc. devops@airhelp.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"

	"github.com/AirHelp/treasury/client"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write ENVIRONMENT/APPLICATION/KEY SECRET",
	Short: "Write secrets into Treasury",
	Long:  `Write sends data into Treasury at the given key (path).`,
	RunE:  write,
}

func init() {
	RootCmd.AddCommand(writeCmd)
	writeCmd.SuggestFor = []string{"put"}
}

func write(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("Missing Key and Value to write.")
	}
	key := args[0]
	value := args[1]

	treasury, err := client.New(treasuryS3, &client.Options{})
	if err != nil {
		return err
	}
	err = treasury.Write(key, value)
	if err != nil {
		return err
	}

	fmt.Println("Success! Data written to: ", key)
	return nil
}
