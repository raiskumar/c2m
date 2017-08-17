package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Make REST call and read response
func GetRestContent(url, user, pass string) []byte {
	//fmt.Println("URL=", url)
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
