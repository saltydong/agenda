package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)


var loginCmd = &cobra.Command{
	Use:   "login -n [UserName] -c [PassWord]",
	Short: "Login to the meeting system.",
	Long: `Using your UserName and PassWord to login Agenda.
	P.S:If the PassWord is right,you can login Agenda and use it. If you forget the PassWord,you must register another one User`,

	PreRun: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile, "[Execute]", log.Ldate|log.Ltime|log.Lshortfile)
		debugLog.Printf("%v\n", os.Args[1:])
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		entity.AgendaStart()
		n, _ := cmd.Flags().GetString("username")
		c, _ := cmd.Flags().GetString("password")

		defer entity.AgendaEnd()
		if entity.UserLogin(n, c) {
			log.Print("Successfully Login")
		} else {
			log.Fatal("Login fail")
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)


	loginCmd.Flags().StringP("username", "n", "", "logged user's username")
	loginCmd.Flags().StringP("password", "c", "", "logged user's password")
}