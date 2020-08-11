package utils

import (
	"encoding/json"
	"fmt"
	"github.com/itbread/lotteryport/datamodels"
	"io/ioutil"
	"net/http"
	"strings"
)

//{
//"state": 0,
//"message": "查询成功",
//"pageCount": 2,
//"countNum": 156,
//"Tflag": 1,
//"result": [{
//"name": "双色球",
//"code": "2018153",
//"detailsLink": "/c/2018-12-30/447336.shtml",
//"videoLink": "",
//"date": "2018-12-30(日)",
//"week": "日",
//"red": "01,07,17,23,25,31",
//"blue": "11",
//"blue2": "",
//"sales": "363663410",
//"poolmoney": "1289528301"
//}]
//}
type SsqResp struct {
	State     int              `json:"state"`
	Message   string           `json:"message"`
	PageNo    int              `json:"pageNo"`    //第几页
	PageCount int              `json:"pageCount"` //总页数
	CountNum  int              `json:"countNum"`  //数目
	Tflag     int              `json:"Tflag"`
	Result    []datamodels.Ssq `json:"result"`
}

type SsqBallInfo struct {
	State     int              `json:"state"`
	Message   string           `json:"message"`
	PageCount int              `json:"pageCount"`
	CountNum  int              `json:"countNum"`
	Tflag     int              `json:"Tflag"`
	Result    []datamodels.Ssq `json:"result"`
}
type SsqBallInfo22 struct {
	State     int    `json:"state"`
	Message   string `json:"message"`
	PageCount int    `json:"pageCount"`
	CountNum  int    `json:"countNum"`
	Tflag     int    `json:"Tflag"`
	Result    []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Date        string `json:"date"`
		Week        string `json:"week"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
	} `json:"result"`
}

func SsqHttpClientGet(issueStart int, issueEnd int, pageNo int, resp *SsqResp) error {
	url := fmt.Sprintf("http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&&issueStart=%v&issueEnd=%v&pageNo=%v", issueStart, issueEnd, pageNo)
	mp := make(map[string]string)
	//增加header选项
	mp["Accept"] = "application/json, text/javascript, */*; q=0.01"
	mp["Accept-Encoding"] = "gzip, deflate"
	mp["Accept-Language"] = "en-US,en;q=0.9"
	mp["Connection"] = "keep-alive"
	//mp["Content-Type"] = "application/json;charset=UTF-8"
	mp["Host"] = "www.cwl.gov.cn"
	mp["Referer"] = "http://www.cwl.gov.cn/kjxx/ssq/kjgg/"
	mp["X-Requested-With"] = "XMLHttpRequest"
	mp["Cookie"] = "_Jo0OQK=290BF4546030AB1D6A14FFE85A88ADA99B9923F1E055FECB1B1EE338B4EEAB2DBDB1DFA5415F8663ABCEB24AC096916E26B03D2654A487FE43802817B6DBA91A1FDB9652B349A2849AA5D71C9A647C60CA55D71C9A647C60CA54F27B4D560638D7FGJ1Z1RQ=="
	return GetRequest(url, mp, resp)
}

func DoHttpget(issueStart int, issueEnd int, pageNo int, resp interface{}) error {

	url := fmt.Sprintf("http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&&issueStart=%v&issueEnd=%v&pageNo=%v", issueStart, issueEnd, pageNo)
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", " application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", " gzip, deflate")
	req.Header.Add("Accept-Language", " en-US,en;q=0.9")
	req.Header.Add("Host", " www.cwl.gov.cn")
	req.Header.Add("Referer", " http://www.cwl.gov.cn/kjxx/ssq/kjgg/")
	req.Header.Add("X-Requested-With", " XMLHttpRequest")
	req.Header.Add("Cookie", "_Jo0OQK=290BF4546030AB1D6A14FFE85A88ADA99B9923F1E055FECB1B1EE338B4EEAB2DBDB1DFA5415F8663ABCEB24AC096916E26B03D2654A487FE43802817B6DBA91A1FDB9652B349A2849AA5D71C9A647C60CA55D71C9A647C60CA54F27B4D560638D7FGJ1Z1RQ==")

	res, err := client.Do(req)
	defer res.Body.Close()
	if body, err := ioutil.ReadAll(res.Body); err != nil {
		err = json.Unmarshal(body, &resp)
		//buffer := new(bytes.Buffer)
		//buffer.Write(body)
		//err = json.NewDecoder(buffer).Decode(&tmpresp)
		if err != nil {

			fmt.Println("%%%%%%%%%%%%%%%json.Unmarshal(body, resp) err= ", err, " tmpresp====", resp)
		}
	} else {
		//fmt.Println(string(body))
	}
	fmt.Println("tmpresp=====", resp)

	return nil
}

func GetSsqHistory(issueStart int, issueEnd int, pageNo int, resp interface{}) error {
	url := fmt.Sprintf("http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&&issueStart=%v&issueEnd=%v&pageNo=%v", issueStart, issueEnd, pageNo)
	method := "GET"
	payload := strings.NewReader("")

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", " application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", " gzip, deflate")
	req.Header.Add("Accept-Language", " en-US,en;q=0.9")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Host", " www.cwl.gov.cn")
	req.Header.Add("Referer", " http://www.cwl.gov.cn/kjxx/ssq/kjgg/")
	req.Header.Add("X-Requested-With", " XMLHttpRequest")
	req.Header.Add("Cookie", "_Jo0OQK=56DADB9BECDA87E0C634D511106103397B94DF68802B86557F473F0C91F14D6003C1D0C785395717B723D546E27D3996097A5BDF5B766BAE0F3BA7DA69DB62E60A7B9652B349A2849AA5D71C9A647C60CA55D71C9A647C60CA54F27B4D560638D7FGJ1Z1OQ==")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &resp)
	fmt.Println("=====================", resp)
	return err
}
