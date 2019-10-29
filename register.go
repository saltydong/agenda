package cmd

import (
	"log"

	"github.com/spf13/cobra"


)


var registerCmd = &cobra.Command{
	Use:   "regist -n [UserName] -c [PassWord] -e [Email] -t [Phone]",
	Short: "register a new user",

	Long: `To create a news user for the meeting agenda system. You have to come up with a username, your password, your email, your phone.
	[UserNme] is new user's name
	[PassWord] is new user's password
	[Email] is new user's email
	[Phone] is new user's phone.`,

	Run: func(cmd *cobra.Command, args []string) {

		entity.ReadFromFile()
		n, _ := cmd.Flags().GetString("username")
		c, _ := cmd.Flags().GetString("password")
		e, _ := cmd.Flags().GetString("email")
		t, _ := cmd.Flags().GetString("phone")

	
		defer entity.AgendaEnd()
		if entity.UserRegister(n, c, e, t) {
			log.Print("Successful Rergister")
		} else {
			log.Fatal("Register fail")
		}

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)


	registerCmd.Flags().StringP("username", "n", "", "new user's username")
	registerCmd.Flags().StringP("password", "c", "", "new user's password")
	registerCmd.Flags().StringP("email", "e", "", "new user's email")
	registerCmd.Flags().StringP("phone", "t", "", "phone new user's phone")


}