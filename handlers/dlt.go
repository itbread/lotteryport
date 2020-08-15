package handlers

import (
	"github.com/itbread/lotteryport/datamodels"
	"github.com/itbread/lotteryport/utils"
	"github.com/kataras/iris/v12"
	"strings"
)

//
func (h *dltHandler) GetDltHistory(ctx iris.Context) {

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

	ctx.JSON(SubErr(Success, errText[Success]))
	return
}
func (h *dltHandler) GetDltList(ctx iris.Context) {

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
	h.service.GetDlts(offset, limit, mp)
	ctx.JSON(SubErr(Success, errText[Success]))
	return
}

func ConvertToDlt(text string) datamodels.Dlt {
	var dlt datamodels.Dlt
	arr := strings.Split(text, "\t")
	if len(arr) > 6 {
		dlt.Code = arr[0]
		dlt.Red1 = arr[1]
		dlt.Red2 = arr[2]
		dlt.Red3 = arr[3]
		dlt.Red4 = arr[4]
		dlt.Red5 = arr[5]
		dlt.Blue1 = arr[6]
		dlt.Blue2 = arr[7]
	}
	return dlt
}

func SaveDlt(path string) {
	lines, err := utils.ReadLine(path)
	if err != nil {
		return
	}

	for _, text := range lines {
		if dlt := ConvertToDlt(text); len(dlt.Code) > 0 {

		}
	}
}

func (h *dltHandler) SaveDltHistory(ctx iris.Context) {
	path := "dlt.txt"
	lines, err := utils.ReadLine(path)
	if err != nil {
		return
	}
	var arr []datamodels.Dlt
	for _, text := range lines {
		if dlt := ConvertToDlt(text); len(dlt.Code) > 0 {
			arr = append(arr, dlt)
		}
	}
	if len(arr) > 0 {
		if err=h.service.Createlist(arr);err!=nil{
			ctx.JSON(SubErr(105, "save to db failed"))
			return
		}
	}

	ctx.JSON(SubErr(Success, errText[Success]))
}
