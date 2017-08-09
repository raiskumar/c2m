package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/raiskumar/c2m/commands"
	"github.com/raiskumar/c2m/vo"
)

// Main method
func main() {
	err := commands.CcbCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Test main method
func mainbkp() {

	//bucket
	contents := commands.GetContent("http://mocky.io/v2/598aa61d410000d51d8211bf", "", "")
	var obj1 []vo.BucketResp
	json.Unmarshal(contents, &obj1)
	fmt.Println("bucket======", obj1[0])
}
