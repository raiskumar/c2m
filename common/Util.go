package common

import (
	"fmt"
	"os"
)

func ValidateCommand(NodeURL string) {
	if len(NodeURL) == 0 {
		fmt.Println("Please configure Your Cluster !")
		fmt.Println("Run config command; for help $./c2m config --help")
		os.Exit(1)
	}
}
