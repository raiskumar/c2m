package commands

import (
	"fmt"
	"io/ioutil"
	"strings"

	b64 "encoding/base64"

	"github.com/spf13/cobra"
)

// RootCmd is the starting point of application
var RootCmd = &cobra.Command{
	Use:   "c2m",
	Short: "c2m is a couchbase cluster manager",
	Long: ` c2m : Couchbase Cluster Manager
Helps in getting key insights of your Couchbase Cluster. 
Ideal for monitoring and managing all the nodes of your Couchbase cluster
Complete documentation available at https://github.com/raiskumar/c2m`,

	// Method will get called if user doesn't provide any command!
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" App name should be followed by Command and then optional sub Command and flags ")
		fmt.Println(" i.e. $./c2m Command sub-Command --verbose")
		fmt.Println(" $./c2m cluster   #to print cluster details!")
		fmt.Println(" Complete documentation available at https://github.com/raiskumar/c2m")
	},
}

//Declare Global variables
var Verbose bool
var NodeURL, UserID, Password, SetupMsg string

// Initializes all available commands
func init() {
	config, err := ioutil.ReadFile("config") // read contents of config file
	if err == nil {
		sDec, _ := b64.StdEncoding.DecodeString(string(config))
		result := strings.Split(string(sDec), ",")
		NodeURL = result[0]
		UserID = result[1]
		Password = result[2]
		//printConfig()
	}

	SetupMsg = " Please configure URL of the cluster !"

	//flags can be used in all commands so it needs to be pre-declared in global scope
	// Verbose variable will be set to true if the format is ./c2m command subcommand --verbose
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(clusterCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.AddCommand(bucketCmd)
	RootCmd.AddCommand(nodeCmd)
	RootCmd.AddCommand(indexCmd)
}

// For debugging
func printConfig() {
	fmt.Println("URL=", NodeURL)
	fmt.Println("User=", UserID)
	fmt.Println("Password=", Password)
}
