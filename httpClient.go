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
func getContent(url, user, pass string) string {
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
	fmt.Println("get:", string(body))

	return string(body)
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
func mainssssss() {
	//getContent("http://mocky.io/v2/5927ca0c1100003e0c6cccf7", "", "")
	ShowCluster()
	//jsonStr = ``
	//jsonStr = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//jsonStr = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//value := gjson.Get(jsonStr, "storageTotals.ram.total")
	//fmt.Println("value = ", value.String())

	//val, _ := jsonparser.GetString([]byte(jsonStr), "alertsSilenceURL")
	//fmt.Println("val =", val)
}
