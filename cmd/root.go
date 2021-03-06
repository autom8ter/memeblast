// Copyright © 2019 NAME HERE Coleman Word
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "memeblast",
	Long: `
------------------------------------------------------------
 __  __  ____  __  __  ____  ____  __      __    ___  ____ 
(  \/  )( ___)(  \/  )( ___)(  _ \(  )    /__\  / __)(_  _)
 )    (  )__)  )    (  )__)  ) _ < )(__  /(__)\ \__ \  )(  
(_/\/\_)(____)(_/\/\_)(____)(____/(____)(__)(__)(___/ (__) 
------------------------------------------------------------ 
Send sms blasts of memes from a randomized set of phonenumbers

(FOR EXPERIMENTATION PURPOSES ONLY)
(NOT FOR USE BY HUMANS)
SMS Laws: https://www.bulksms.com/resources/regulations/requirements-for-marketing-using-sms-messaging-in-the-usa.htm
------------------------------------------------------------
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigFile("./memeblast.yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(`expected memeblast.yaml containing "from" numbers(twilio) and "to" numbers(targets) in $PWD`, err.Error())
	} else {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	}
}
