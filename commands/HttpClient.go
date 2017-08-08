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

/*type CouchbaseResponse struct {

}*/

//test
func main() {
	//ShowCluster()
	contents := GetContent("http://mocky.io/v2/5986c32d1100009c00fcbe4a", "", "")
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)
	fmt.Println("uri======", obj.Nodes[0].Uptime)

}
