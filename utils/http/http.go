package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get()  {
	client := &http.Client{}
	//生成要访问的url
	url:="http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=2017-10-24&dayEnd=2018-10-24"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Accept","application/json, text/javascript, */*; q=0.01")
	reqest.Header.Add("Content-Type","application/json;charset=UTF-8")
	reqest.Header.Add("Referer","http://www.cwl.gov.cn/kjxx/ssq/kjgg/")
	reqest.Header.Add("Host", "www.cwl.gov.cn")
	reqest.Header.Add("X-Requested-With", "XMLHttpRequest")
	reqest.Header.Add("Cookie", "ites=_21; UniqueID=8Z8HyjtRT9W3373K1560780231892; _ga=GA1.3.95489882.1557983961; FSSBBIl1UgzbN7N80S=LU5FY9X1OFBXK.NNdKw68ItKEWTtMYLk8MbQQz2Yc7mytIhesHDWQclTF0iH0xHF; _gid=GA1.3.625038784.1560775609; 21_vq=24; _gat_gtag_UA_113065506_1=1; FSSBBIl1UgzbN7N80T=3rNx0eiaXUXjUTCYjSiS1hiulKMxLU88Pwc3zvWVR7n7oqXppcgZLKtpBYzDTjOCXdMZX1VaaCeng3OIrvKG4b85r63pRtmYR4RYW7fiEBWzz1KeH2LE0aPDIqJc5LDtjQb2PTza6.P4hkQU2O4jL97Doh7X4B9XhNpzpBLO8_igRnrvy85ZIg8OgzeKpC9W7OqYKCpwcIjn4Fzq2PdNXsWEk7Bz6ZubBVDBn5uYnoDr4BgffTx8SxLkNhIFh39kwXPt5qYFu8PMz76.x0aZplf2gfrWw177y4QKWGCqoFNy.uCIA6EGBCvisOGjUG3x.BXG")
	reqest.Header.Add("Cache-Control", "0")
	reqest.Header.Add("Connection", "keep-alive")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
// http.Get
func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencode",
		strings.NewReader("name=abc")) // Content-Type post请求必须设置
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func HttpDo() {
	client := &http.Client{}
	//http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=2017-10-24&dayEnd=2018-10-24
	url:="http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=2017-10-24&dayEnd=2018-10-24"
	req, err := http.NewRequest("GET", url, strings.NewReader("name=ssq"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer","http://www.cwl.gov.cn/kjxx/ssq/kjgg/")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}