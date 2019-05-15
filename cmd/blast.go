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
	"github.com/autom8ter/smsdos/cats"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// blastCmd represents the blast command
var blastCmd = &cobra.Command{
	Use:   "cats",
	Short: "start an cats sms blast",
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetString("twilio_account") == "" {
			fmt.Println(`expected "twilio_account" containing your twilio account number  in smsdos.yaml file`)
			os.Exit(1)
		}
		if viper.GetString("twilio_key") == "" {
			fmt.Println(`expected "twilio_key" containing your twilio api key in smsdos.yaml file`)
			os.Exit(1)
		}
		if len(viper.GetStringSlice("from")) == 0 {
			fmt.Println(`expected "from" containing an array of your twilio phone numbers in smsdos.yaml file`)
			os.Exit(1)
		}
		if len(viper.GetStringSlice("to")) == 0 {
			fmt.Println(`expected "to" containing an array of phone numbers in smsdos.yaml file`)
			os.Exit(1)
		}
		if viper.GetInt("send") == 0 {
			fmt.Println(`"send" not found int smsdos.yaml`, "setting default send to 1")
			viper.Set("send", 1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		send := viper.GetInt("send")
		for x := 0; x < send; x++ {
			cats.Blast().Attack()
		}
	},
}

func init() {
	rootCmd.AddCommand(blastCmd)
}
