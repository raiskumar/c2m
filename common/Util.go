package common

import (
	"fmt"
	"os"
	"strconv"
)

func ValidateCommand(NodeURL string) {
	if len(NodeURL) == 0 {
		fmt.Println("Please configure Your Cluster !")
		fmt.Println("Run config command; for help $./c2m config --help")
		os.Exit(1)
	}
}

// Converts byte into human redable format
func HumanRedableMemory(val int64) string {
	tmp := val >> 30

	if tmp > 0 {
		return strconv.FormatInt(tmp, 10) + " GB"
	}

	tmp = val >> 20
	if tmp > 0 {
		return strconv.FormatInt(tmp, 10) + " MB"
	}

	tmp = val >> 10
	return strconv.FormatInt(tmp, 10) + " KB"
}

func HumandRedableUpFor(val string) string {
	//fmt.Println(" up for " + val)
	int64Val, error := strconv.ParseInt(val, 10, 64)
	if error != nil {
		return val
	}

	mins := int64Val / 60

	if mins < 60 {
		return strconv.FormatInt(mins, 10) + " mins"
	}

	hours := mins / 60

	if hours < 24 {
		return strconv.FormatInt(hours, 10) + " hours & " + strconv.FormatInt(mins%60, 10) + " mins"
	}

	days := hours / 24
	return strconv.FormatInt(days, 10) + " day(s) & " + strconv.FormatInt(hours%60, 10) + " hours"
}
