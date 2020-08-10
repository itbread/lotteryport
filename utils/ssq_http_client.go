package utils

import (
	"fmt"
	"github.com/itbread/lotteryport/datamodels"
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

	return  GetRequest(url, mp, resp)

}
