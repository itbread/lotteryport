package handlers

import (
	"net/http"
	"strconv"
)

const Success = 0             //成功
const Unauthorized = 4001     //鉴权末通过
const PermissionDenied = 4002 //权限不足
const WeiXinSessionErr = 4003 //微信小程序session超时或无效
const UserSessionErr = 4004   //用户session超时或无效
const ParamsErr = 5001        //参数错误
const ParamsMiss = 5002       //参数缺失
const OperationFailed = 6001  //业务处理失败
const OperationTimeout = 6002 //业务处理超时

var errText = map[int]string{
	Success:          "Success",                           //成功
	Unauthorized:     "Unauthorized",                      //鉴权末通过
	PermissionDenied: "Permission Denied",                 //权限不足
	WeiXinSessionErr: "WeiXin Session Timeout Or Invalid", //微信小程序session超时或无效
	UserSessionErr:   "User Session Timeout Or Invalid",   //用户session超时或无效
	ParamsErr:        "Params Error",                      //参数错误
	ParamsMiss:       "Params Missing",                    //参数缺失
	OperationFailed:  "Operation Failed",                  //业务处理失败
	OperationTimeout: "Operation Timeout",                 //业务处理超时
}

func ErrText(code int) string {
	return errText[code]
}

// 通用返回
type Resp struct {
	Code    int         `json:"code"`           // 通用错误码
	Msg     string      `json:"msg"`            // 通用错误描述
	SubCode int         `json:"sub_code"`       // 业务错误码
	SubMsg  string      `json:"sub_msg"`        // 业务错误描述
	Data    interface{} `json:"data,omitempty"` // 业务数据
}

// 给返回加业务成功
func SubSuccess(data ...interface{}) (resp *Resp) {
	subCode := Success
	subMsg := ErrText(Success)
	switch len(data) {
	case 1:
		return SubErr(subCode, subMsg, data[0])
	default:
		return SubErr(subCode, subMsg)
	}
}

// 给返回加业务错误码
func SubErr(subCode int, subMsg string, data ...interface{}) (resp *Resp) {
	switch len(data) {
	case 1:
		resp = &Resp{Code: Success,
			Msg:     ErrText(Success),
			SubCode: subCode,
			SubMsg:  subMsg,
			Data:    data[0],
		}
		return
	default:
		resp = &Resp{Code: Success,
			Msg:     ErrText(Success),
			SubCode: subCode,
			SubMsg:  subMsg,
		}
		return
	}
}

// 给返回加错误码
func Err(Code int) *Resp {
	return &Resp{Code: Code, Msg: ErrText(Code)}
}

const (
	DefaultLimitSize = 10 //默认Limit大小
)

//读取*http.Request请求中的url 参数
func ReadOffsetAndLimit(r *http.Request) (offset int, limit int) {
	var err error
	offset, err = strconv.Atoi(r.FormValue("offset"))
	if err != nil || offset == 0 {
		offset = 0
	}
	limit, err = strconv.Atoi(r.FormValue("limit"))
	if err != nil || limit == 0 {
		limit = DefaultLimitSize
	}
	return offset, limit
}

//读取*http.Request请求中的url int64参数
func GetInt64FromRequest(r *http.Request, key string, defaultvalue int64) int64 {
	value, err := strconv.ParseInt(r.FormValue(key), 10, 64)
	if err != nil {
		value = defaultvalue
	}
	return value
}

//读取*http.Request请求中的url uint64参数
func GetUInt64FromRequest(r *http.Request, key string, defaultvalue uint64) uint64 {
	value, err := strconv.ParseUint(r.FormValue(key), 10, 64)
	if err != nil {
		value = defaultvalue
	}
	return value
}

//读取*http.Request请求中的int参数
func GetIntFromRequest(r *http.Request, key string, defaultvalue int) int {
	value, err := strconv.Atoi(r.FormValue(key))
	if err != nil {
		value = defaultvalue
	}
	return value
}
