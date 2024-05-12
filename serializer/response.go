package serializer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ResponseNotWithData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// swagger:response Response
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

const (
	CodeCheckLogin = 401
	CodeNoRightErr = 403

	CodeCommonError  = 50000
	CodeDBError      = 50001
	CodeEncryptError = 50002
	CodeParamErr     = 50003
	SignExpired      = 50004
	AlreadyMinted    = 50005
	NotWhitelist     = 50006
	NftMintOver      = 50007
	SignVerifyFailed = 50008

	PaySuccessed = 50009

	DATA_NOT_FOUND_IN_NET = 50011
	LOWER_OVER_FEE        = 50012
	CONFIG_ERR            = 50013
	TIME_PASER_ERR        = 50015
	REQUEST_NET_DATA_ERR  = 50016
)

func Fail(errCode int, msg string) Response {
	return Response{
		Code: errCode,
		Msg:  msg,
	}
}

func FailWithData(errCode int, data interface{}, msg string) Response {
	return Response{
		Code: errCode,
		Data: data,
		Msg:  msg,
	}
}

func WriteData2Front(c *gin.Context, resData interface{}, err error, defineMsg string) {
	if err != nil {
		WrapErrReturn(c, defineMsg, err)
		return
	}
	c.JSON(http.StatusOK, Success(resData))
}

func Success(data interface{}) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
	}
}

func SuccessNotWithData() ResponseNotWithData {
	return ResponseNotWithData{
		Code: http.StatusOK,
	}
}

func WrapErrReturn(c *gin.Context, defineMsg string, err error) {
	if err != nil {
		log.Println(err.Error(), err)
		if defineMsg == "" {
			defineMsg = err.Error()
		}
	}
	c.JSON(200, Fail(http.StatusInternalServerError, defineMsg))
}
