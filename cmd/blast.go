// Copyright © 2019 Coleman Word
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
	"fmt"
	"github.com/autom8ter/smsdos/messages"
	"github.com/spf13/cobra"
	"os"
)

var blast *messages.Blast

// blastCmd represents the blast command
var blastCmd = &cobra.Command{
	Use:   "blast",
	Short: "start an sms blast (default: cat memes)",
	PreRun: func(cmd *cobra.Command, args []string) {
		blast = messages.NewBlast()
		if err := blast.Validate();err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		blast.Attack()
	},
}

func init() {
	rootCmd.AddCommand(blastCmd)
}
