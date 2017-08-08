package commands

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

// ./c2m node
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Pass cluster details - helps in bootstraping the tool",
	Long: `Command which takes cluster details from user; 
            uri, user, pass
			$./c2m config <URI>
			or
			$./c2m config <URI> <USER> <PASS>
            `,
	Run: config,
}

// config command
func config(cmd *cobra.Command, args []string) {
	URI := "http://mocky.io/v2/5986c32d1100009c00fcbe4a"
	USER := ""
	PASS := ""

	// Validation check
	// Minimum one parameter should be passed

	for i, val := range args {
		if i == 0 {
			URI = val
		}
		if i == 1 {
			USER = val
		}
		if i == 2 {
			PASS = val
		}
	}
	os.Setenv("URI", URI)
	os.Setenv("USER", USER)
	os.Setenv("PASS", PASS)
	fmt.Println("URI=", os.Getenv("URI"))
	//fmt.Println("USER=", os.Getenv("USER"))
	//fmt.Println("PASS=", os.Getenv("PASS"))

	env := os.Environ()
	env = append(env, fmt.Sprintf("URI=%s", URI))

	fmt.Println("List of Environtment variables : \n")

	/*for index, value := range env {
		name := strings.Split(value, "=") // split by = sign

		fmt.Printf("[%d] %s : %v\n", index, name[0], name[1])
	}*/
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
