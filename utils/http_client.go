package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	DELETE_METHOD = "DELETE"
)

//发送GET请求,有两种请求方式
// Get(url , respPtr)
// Get(url , reqPtr, respPtr)
// url: 请求链接
// reqPtr: 请求数据的结构体地址
// respPtr： 返回数据的结构体地址
func Get(url string, args ...interface{}) error {
	log.Printf("Send Get Request, Url:%v\n", url)
	mp := make(map[string]string)
	return request(GET_METHOD, url, mp, args...)
}

func GetRequest(url string, mp map[string]string, args ...interface{}) error {
	log.Printf("Send Get Request, Url:%v\n", url)
	return request(GET_METHOD, url, mp, args...)
}

//发送DELETE请求,有两种请求方式
// Delete(url , respPtr)
// Delete(url , reqPtr, respPtr)
// url: 请求链接
// reqPtr: 请求数据的结构体地址
// respPtr： 返回数据的结构体地址
func Delete(url string, args ...interface{}) error {
	log.Printf("Send Delete Request, Url:%v\n", url)
	mp := make(map[string]string)
	return request(DELETE_METHOD, url, mp, args...)
}
func DeleteRequest(url string, mp map[string]string, args ...interface{}) error {
	log.Printf("Send Delete Request, Url:%v\n", url)
	return request(DELETE_METHOD, url, mp, args...)
}

//发送POST请求,有两种请求方式
// Post(url , respPtr)
// Post(url , reqPtr, respPtr)
// url: 请求链接
// reqPtr: 请求数据的结构体地址
// respPtr： 返回数据的结构体地址
func Post(url string, args ...interface{}) error {
	log.Printf("Send Post Request, Url:%v\n", url)
	mp := make(map[string]string)
	return request(POST_METHOD, url, mp, args...)
}

func request(method string, url string, head map[string]string, args ...interface{}) error {

	var (
		reqStruct  interface{}
		respStruct interface{}
		req        *http.Request
		resp       *http.Response
		client     http.Client
		data       string
		respData   []byte
		err        error
	)

	switch len(args) {
	case 1:
		respStruct = args[0]
	case 2:
		reqStruct = args[0]
		respStruct = args[1]
	default:
		return errors.New("function must have 2 or 3 arguments")
	}
	sendBody, jsonErr := json.Marshal(reqStruct)
	if  jsonErr != nil {
		return jsonErr
	} else {
		data = string(sendBody)
	}
	log.Printf("Send Data:%v\n", data)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err = http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if len(head) > 0 {
		for k, v := range head {
			req.Header.Add(k, v)
		}
	} else {
		req.Header.Add("Content-Type", "application/json; charset=utf-8") //设置content-type
	}
	req.Header.Add("Content-Length", fmt.Sprintf("%v",len(sendBody))) //设置content-type
	fmt.Printf("======v%",req.Header)
	resp, err = client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%v", resp.StatusCode))
	}
	respData, err = ioutil.ReadAll(resp.Body)
	log.Printf("Response Data:%s\n", respData)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	buffer.Write(respData)
	err = json.NewDecoder(buffer).Decode(&respStruct)
	if err != nil {
		return err
	}
	return nil
}
func httpRequest(method string, url string, head map[string]string, args ...interface{}) error {

	var (
		reqStruct  interface{}
		respStruct interface{}
		req        *http.Request
		resp       *http.Response
		client     http.Client
		data       string
		respData   []byte
		err        error
	)

	switch len(args) {
	case 1:
		respStruct = args[0]
	case 2:
		reqStruct = args[0]
		respStruct = args[1]
	default:
		return errors.New("function must have 2 or 3 arguments")
	}
	if sendBody, jsonErr := json.Marshal(reqStruct); jsonErr != nil {
		return jsonErr
	} else {
		data = string(sendBody)
	}

	log.Printf("Send Data:%v\n", data)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err = http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if len(head) > 0 {
		for k, v := range head {
			req.Header.Add(k, v)
		}
	} else {
		req.Header.Add("Content-Type", "application/json; charset=utf-8") //设置content-type
	}
	resp, err = client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%v", resp.StatusCode))
	}
	respData, err = ioutil.ReadAll(resp.Body)
	log.Printf("Response Data:%s\n", respData)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	buffer.Write(respData)
	err = json.NewDecoder(buffer).Decode(&respStruct)
	if err != nil {
		return err
	}
	return nil
}
