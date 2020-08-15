package handlers

import (
	"github.com/itbread/lotteryport/datamodels"
	"github.com/itbread/lotteryport/utils"
	"github.com/kataras/iris/v12"
	"log"
	"strings"
)

//
func (h *ssqHandler) GetSsqHistory(ctx iris.Context) {

	const (
		Success               = 0   // 成功
		OpearteErr            = 101 // 操作失败
		RequstUserServiceErr  = 102
		VerifyCodeErr         = 103 // 验证码不正确
		OperationErr          = 104
		ReasonIdEmptyErr      = 105
		WorkerCanNotSubmitErr = 106
		OtherErr              = 301
	)
	var errText = map[int]string{
		Success:               "Success", //成功
		OpearteErr:            "Phone Number Format Error",
		VerifyCodeErr:         "Wrong Code",
		OperationErr:          "Operation Error",
		RequstUserServiceErr:  "Request User Service Error",
		ReasonIdEmptyErr:      "ReasonId Empty  Error",
		WorkerCanNotSubmitErr: "Worker Can Not GetSsqHistory Error",
		OtherErr:              "Other Errors",
	}
	var resp utils.SsqResp
	issueStart := int(ctx.URLParamInt32Default("issueStart", 0))
	issueEnd := int(ctx.URLParamInt32Default("issueEnd", 0))
	pageNo := int(ctx.URLParamInt32Default("pageNo", 1))
	//offset, limit := ReadOffsetAndLimit(ctx.Request())
	//mp := make(map[string]interface{})
	//h.service.GetDlts(offset, limit, mp)
	//err := utils.SsqHttpClientGet(issueStart, issueEnd, pageNo, &resp)
	err := utils.GetSsqHistory(issueStart, issueEnd, pageNo, &resp)
	if err != nil {
		log.Printf("请求 ssq服务器.Error: %v\n", err)
		ctx.JSON(SubErr(OpearteErr, errText[OpearteErr], resp))
		return
	} else {
		log.Printf("存储数据到数据库:")
		var list    []datamodels.Ssq
		if len(resp.Result) > 0 {
			for _,item:=range resp.Result{
				if arr:=strings.Split(item.Red,",");len(arr)>5{
					item.Red1=arr[0]
					item.Red2=arr[1]
					item.Red3=arr[2]
					item.Red4=arr[3]
					item.Red5=arr[4]
					item.Red6=arr[5]
				}
				list=append(list,item)
			}
			log.Printf("==========存储数据到数据库:")
			err = h.service.Createlist(list)
			if err != nil {
				log.Printf("========Createlist.Error: %v\n", err)
				ctx.JSON(SubErr(OpearteErr, errText[OpearteErr], resp))
				return
			}
		} else {
			log.Printf("resp.Result=====:%v", resp.Result)
		}

	}
	ctx.JSON(SubErr(Success, errText[Success], resp))
	return
}
func (h *ssqHandler) GetSsqList(ctx iris.Context) {

	const (
		Success               = 0   // 成功
		PhoneFormatErr        = 101 // 手机号码格式不正确
		RequstUserServiceErr  = 102
		VerifyCodeErr         = 103 // 验证码不正确
		OperationErr          = 104
		ReasonIdEmptyErr      = 105
		WorkerCanNotSubmitErr = 106
		OtherErr              = 301
	)
	var errText = map[int]string{
		Success:               "Success", //成功
		PhoneFormatErr:        "Phone Number Format Error",
		VerifyCodeErr:         "Wrong Code",
		OperationErr:          "Operation Error",
		RequstUserServiceErr:  "Request User Service Error",
		ReasonIdEmptyErr:      "ReasonId Empty  Error",
		WorkerCanNotSubmitErr: "Worker Can Not GetSsqHistory Error",
		OtherErr:              "Other Errors",
	}
	//var resp services.SsqResp
	//issueStart := int(ctx.URLParamInt32Default("issueStart", 0))
	//issueEnd := int(ctx.URLParamInt32Default("issueEnd", 0))
	//pageNo := int(ctx.URLParamInt32Default("pageNo", 1))
	offset, limit := ReadOffsetAndLimit(ctx.Request())
	mp := make(map[string]interface{})
	h.service.GetSsqs(offset, limit, mp)
	ctx.JSON(SubErr(Success, errText[Success]))
	return
}
