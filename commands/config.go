package commands

import (
	"fmt"
	"os"

	b64 "encoding/base64"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Bootstraps the CLI by configuring the Url and Access Credentials!",
	Long: `Command which takes uri, and optional user-id, password of any node
$./c2m config <URI>
or
$./c2m config <URI> <USER> <PASS> `,
	Run: config,
}

// config command
func config(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Node URL is mandatory!")
		os.Exit(1)
	}

	if len(args) == 3 {
		UserID = args[1]
		Password = args[2]
	}
	NodeURL = args[0]

	file, err := os.Create("config")
	if err != nil {
		fmt.Println("Not able to create file -", err)
		return
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(NodeURL + "," + UserID + "," + Password))
	file.Write([]byte(sEnc))
}
