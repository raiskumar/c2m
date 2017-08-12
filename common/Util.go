package common

import (
	"fmt"
	"os"
)

func ValidateCommand(NodeURL string) {
	if len(NodeURL) == 0 {
		fmt.Println(" Please configure URL of the cluster !")
		os.Exit(1)
	}
}
