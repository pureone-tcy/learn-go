package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// http
func main() {
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	base, err := url.Parse("http://example.com")
	fmt.Println(base, err)

	ref, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(ref).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `w/xyzzy`)
	q := req.URL.Query()
	fmt.Println(q)

	q.Add("c", "3&%")
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// POST
	//req2, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("foo")))
}
