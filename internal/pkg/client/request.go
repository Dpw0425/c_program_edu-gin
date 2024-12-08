package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestClient struct {
	runMode string
	client  *http.Client
}

func NewRequestClient(client *http.Client) *RequestClient {
	return &RequestClient{
		client: client,
	}
}

func (r *RequestClient) SetDebug() {
	r.runMode = "debug"
}

func (r *RequestClient) Get(url string, params *url.Values) ([]byte, error) {
	if params != nil {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, params.Encode())
		} else {
			url = fmt.Sprintf("%s?%s", url, params.Encode())
		}
	}

	resp, err1 := r.client.Get(url)
	if err1 != nil {
		return nil, err1
	}
	defer resp.Body.Close()

	btResp, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	if r.runMode == "debug" {
		fmt.Printf("\n[GET] HTTP Request\n")
		fmt.Printf("Request URL : %s\n", url)
		fmt.Printf("NewResponse StatusCode: %d\n", resp.StatusCode)
		fmt.Printf("NewResponse Data: %s\n\n", string(btResp))
	}

	return btResp, nil
}

func (r *RequestClient) PostForm(url string, params *url.Values) ([]byte, error) {
	resp, err1 := r.client.PostForm(url, *params)
	if err1 != nil {
		return nil, err1
	}
	defer resp.Body.Close()

	btResp, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return btResp, err2
	}

	if r.runMode == "debug" {
		fmt.Printf("\n[POST] HTTP Request\n")
		fmt.Printf("Request URL : %s\n", url)
		fmt.Printf("Request Data: %s\n", params.Encode())
		fmt.Printf("NewResponse StatusCode: %d\n", resp.StatusCode)
		fmt.Printf("NewResponse Data: %s\n\n", string(btResp))
	}

	return btResp, nil
}

func (r *RequestClient) PostJSON(url string, params *url.Values) ([]byte, error) {
	text, _ := json.Marshal(params)

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(text)))

	req.Header.Set("Content-Type", "application/json")

	resp, err1 := r.client.Do(req)
	if err1 != nil {
		return nil, err1
	}
	defer resp.Body.Close()

	btResp, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	if r.runMode == "debug" {
		fmt.Printf("\n[POST] HTTP Request\n")
		fmt.Printf("Request URL : %s\n", url)
		fmt.Printf("Request Data: %s\n", string(text))
		fmt.Printf("NewResponse StatusCode: %d\n", resp.StatusCode)
		fmt.Printf("NewResponse Data: %s\n\n", string(btResp))
	}

	return btResp, nil
}
