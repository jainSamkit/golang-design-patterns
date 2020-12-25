package network

import (
	"errors"
	"fmt"
	"net/http"
)

//Network strcuture
type Network struct {
	Headers map[string]string
}

//SetHeader sets header field
func (n *Network) SetHeader(key, val string) {
	if (n.Headers) == nil {
		n.Headers = make(map[string]string)
	}
	n.Headers[key] = val
}

//Get request gets response from the url
func (n *Network) Get(url string) (*http.Response, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	for key, val := range n.Headers {
		req.Header.Set(key, val)
	}
	res, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error occurred GET Request: %v", err)
		return res, errors.New(errMsg)
	}
	return res, err
}
