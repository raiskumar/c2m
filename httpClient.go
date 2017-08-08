package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/raiskumar/c2m/vo"
)

// url should be fully formed with port as well (like http://172.17.0.4:8091)
func GetContent(url, user, pass string) []byte {
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

	return body
}

func test() string {
	response, err := http.Get("http://mocky.io/v2/5986c32d1100009c00fcbe4a")
	//5927ca0c1100003e0c6cccf7
	var contents string
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("------%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", contents)
	}
	//return string(contents)
	return contents
}

var objmap vo.PoolResp

func ShowCluster() string {
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
		//fmt.Printf("%s\n", string(contents))
		res := string(contents)
		fmt.Println(res)

		json.Unmarshal(contents, &objmap)
		fmt.Println("uri======", objmap.Nodes[0].Uptime)
	}

	return res
}

/*type CouchbaseResponse struct {

}*/

//test
func mainss() {
	//ShowCluster()
	contents := GetContent("http://mocky.io/v2/5986c32d1100009c00fcbe4a", "", "")
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)
	fmt.Println("uri======", obj.Nodes[0].Uptime)

}
