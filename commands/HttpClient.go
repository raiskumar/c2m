package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/raiskumar/c2m/vo"
)

// url should be fully formed with port as well (like http://172.17.0.4:8091)
func GetContent(url, user, pass string) []byte {
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

//test
func mainsss() {
	contents := GetContent("http://mocky.io/v2/5986c32d1100009c00fcbe4a", "", "")
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)
	fmt.Println("uri======", obj.Nodes[0].Uptime)

	//bucket
	contents = GetContent("http://mocky.io/v2/598aa61d410000d51d8211bf", "", "")
	var obj1 []vo.BucketResp
	json.Unmarshal(contents, &obj1)
	fmt.Println("bucket======", obj1[0])
}
