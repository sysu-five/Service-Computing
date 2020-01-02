/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/sysu-five/agenda/src"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [telephone]",
	Short: "Register a new user",
	Long: `This command is used like: register -n Test -p 123456 -e 123456789@qq.com -t 15300000000`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		u_name, _ := cmd.Flags().GetString("name")
		u_password, _ := cmd.Flags().GetString("password")
		u_email, _ := cmd.Flags().GetString("email")
		u_telephone, _ := cmd.Flags().GetString("telephone")
		if src.IsLogin() == false {
			src.RegisterUser(u_name,u_password,u_email,u_telephone)
		} else {
			fmt.Println("You've already logged in, please log out first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("name","n","","user name")
	registerCmd.Flags().StringP("password","p","","user password")
	registerCmd.Flags().StringP("email","e","","user email")
	registerCmd.Flags().StringP("telephone","t","","user telephone")
}
