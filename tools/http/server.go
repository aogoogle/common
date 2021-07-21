package http

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Headers map[string]string

type Cookies []*http.Cookie

type Data map[string]interface{}

type Params struct {
	Data Data
	Headers Headers
	Cookies Cookies
}

var Debugger bool = false

type Responses struct {
	Response *http.Response
	Body string
}

func request(method string, url string, params Params) (r Responses, e error) {
	// 打印debugger发起请求
	printSendDebugger(method, url, params)
	r, e = do(method, url, params)
	if e != nil {
		return
	}
	// 打印debugger返回请求
	printResultDebugger(r)
	return
}

func do(method string, url string, params Params) (r Responses, e error) {
	request, e := http.NewRequest(method, url, nil)
	// header
	addHeaders(request, params.Headers)
	// cookie
	addCookies(request, params.Cookies)
	// 参数
	addData(request, params.Data)
	r.Response, e = http.DefaultClient.Do(request)
	defer r.Response.Body.Close()
	if e != nil {
		return
	}
	unCoding(&r)
	return
}

func addHeaders(request *http.Request, headers Headers) {
	for k, v := range headers{
		request.Header.Add(k, v)
	}
}

// 添加cookie
func addCookies(request *http.Request, cookies Cookies) {
	for _, v := range cookies{
		request.AddCookie(v)
	}
}

// 添加参数
func addData(request *http.Request, data Data) {
	query := request.URL.Query()
	for k, v := range data{
		query.Add(k, fmt.Sprint(v))
	}
	request.URL.RawQuery = query.Encode()
}

func unCoding(r *Responses) {
	if r.Response.StatusCode == 200 {
		switch r.Response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(r.Response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				r.Body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(r.Response.Body)
			r.Body = string(bodyByte)
		}
	} else {
		bodyByte, _ := ioutil.ReadAll(r.Response.Body)
		r.Body = string(bodyByte)
	}
}

// debugger发起请求
func printSendDebugger(method string, url string, params Params) {
	if Debugger {
		log.Println("debug log start ----------")
		fmt.Println("Method", method)
		fmt.Println("Host", ":", url)
		for k, v := range params.Headers{
			fmt.Println(k, ":", v)
		}
		fmt.Println("----------------------------------------------------")
	}
}

func printResultDebugger(r Responses) {
	if Debugger {
		fmt.Println("Status", ":", r.Response.Status)
		for key, val := range r.Response.Header {
			fmt.Println(key, ":", val[0])
		}
		log.Println("debug log end ----------")
	}
}

func Request(method string, url string, params Params) (Responses, error) {
	return request(method, url , params)
}

func Get(url string, params Params) (Responses, error) {
	return request("GET", url , params)
}

func Post(url string, params Params) (Responses, error) {
	return request("POST", url , params)
}

func Put(url string, params Params) (Responses, error) {
	return request("PUT", url , params)
}

func Delete(url string, params Params) (Responses, error) {
	return request("DELETE", url , params)
}
