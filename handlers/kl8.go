package handlers

import (
	"github.com/itbread/lotteryport/datamodels"
	"github.com/itbread/lotteryport/utils"
	"github.com/kataras/iris/v12"
	"log"
	"strings"
)

//
func (h *kl8Handler) GetKl8History(ctx iris.Context) {

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
	var resp utils.Kl8Resp
	issueStart := int(ctx.URLParamInt32Default("issueStart", 0))
	issueEnd := int(ctx.URLParamInt32Default("issueEnd", 0))
	pageNo := int(ctx.URLParamInt32Default("pageNo", 1))
	//offset, limit := ReadOffsetAndLimit(ctx.Request())
	//mp := make(map[string]interface{})
	//h.service.GetDlts(offset, limit, mp)
	//err := utils.SsqHttpClientGet(issueStart, issueEnd, pageNo, &resp)
	err := utils.GetKl8History(issueStart, issueEnd, pageNo, &resp)
	if err != nil {
		log.Printf("请求 ssq服务器.Error: %v\n", err)
		ctx.JSON(SubErr(OpearteErr, errText[OpearteErr], resp))
		return
	} else {
		log.Printf("存储数据到数据库:")
		var list []datamodels.Kl8
		if len(resp.Result) > 0 {
			for _, item := range resp.Result {
				if arr := strings.Split(item.Red, ","); len(arr) > 19 {
					item.Red1 = arr[0]
					item.Red2 = arr[1]
					item.Red3 = arr[2]
					item.Red4 = arr[3]
					item.Red5 = arr[4]
					item.Red6 = arr[5]
					item.Red7 = arr[6]
					item.Red8 = arr[7]
					item.Red9 = arr[8]
					item.Red10 = arr[9]
					item.Red11 = arr[10]
					item.Red12 = arr[11]
					item.Red13 = arr[12]
					item.Red14 = arr[13]
					item.Red15 = arr[14]
					item.Red16 = arr[15]
					item.Red17 = arr[16]
					item.Red18 = arr[17]
					item.Red19 = arr[18]
					item.Red20 = arr[19]
				}
				list = append(list, item)
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
func (h *kl8Handler) GetKl8List(ctx iris.Context) {

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
	h.service.GetKl8s(offset, limit, mp)
	ctx.JSON(SubErr(Success, errText[Success]))
	return
}
