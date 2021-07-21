package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HttpGet 发送GET请求
// url：请求地址
func HttpGet(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// HttpPost 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func HttpPost(url string, data interface{}, contentType string) string {
	client := &http.Client{Timeout: 5 * time.Second}

	jsonStr, _ := json.Marshal(data)
	var resp *http.Response
	var err error
	if contentType == "application/x-www-form-urlencoded" {
		resp, err = client.Post(url, contentType, strings.NewReader(data.(string)))
	} else {
		resp, err = client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	}

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func HttpPut(url string, data interface{}, headers map[string]interface{}) string {
	req, _ := http.NewRequest("PUT", url, strings.NewReader(data.(string)))
	//req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value.(string))
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}
	result, _ := ioutil.ReadAll(response.Body)
	return string(result)
}

