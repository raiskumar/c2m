package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/raiskumar/c2m/vo"
)

// Make REST call and read response
func GetRestContent(url, user, pass string) []byte {
	fmt.Println("URL=", url)
	client := &http.Client{}

	/* Authenticate */
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error occured while reading :%s", err)
	}
	return body
}

func TestGet() string {
	url := "http://mocky.io/v2/5986c32d1100009c00fcbe4a"
	var res string
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		res := string(contents)
		fmt.Println(res)
	}
	return res
}

func mainss() {
	contents := GetRestContent("http://mocky.io/v2/5986c32d1100009c00fcbe4a", "", "")
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)
	fmt.Println("uri======", obj.Nodes[0].Uptime)
}
