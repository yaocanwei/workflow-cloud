package handlers

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Pagination struct {
	Count int `json:"count"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

var (
	InternalServerError = errors.New("服务器内部错误")
	LimitError          = errors.New("请输入正确的limit")
	SkipError           = errors.New("请输入正确的skip")
	JWTPrefix           = ""
)

// errorData
func errorData(code int) (data string) {
	codeError := Errors[code]
	if codeError != "" {
		data = codeError
	}
	return
}

// JSONResponse 处理结果集为json格式
func JSONResponse(c *gin.Context, code int, data interface{}, msg string) {
	if data == nil {
		codeError := Errors[code]
		if codeError != "" {
			data = codeError
		}
	}
	c.JSON(200, gin.H{"code": code, "data": data, "msg": msg, "server_time": time.Now()})
}

func ParameterValidateResponse(c *gin.Context, code int, err interface{}, msg string) {
	c.JSON(200, gin.H{"code": code, "error": err, "msg": msg, "server_time": time.Now()})
}

func WsJSONResponse(ws *websocket.Conn, data interface{}, msg string) error {
	type resp struct {
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}

	rs := &resp{}
	rs.Data = data
	rs.Msg = msg

	b, err := json.Marshal(rs)
	if err != nil {
		WsJSONResponse(ws, nil, err.Error())
		return err
	}

	err = ws.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}
	return nil
}

// ParseLimitSkip 格式化分页参数
func ParseLimitSkip(c *gin.Context) (limit int, skip int, err error) {
	limitStr := c.Query("limit")
	skipStr := c.Query("skip")

	if limitStr == "" {
		limitStr = "10"
	}
	if skipStr == "" {
		skipStr = "0"
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		err = LimitError
		return
	}
	skip, err = strconv.Atoi(skipStr)
	if err != nil {
		err = SkipError
		return
	}
	return
}

// JSONResponseWithPagination is response a json, must have status code, data, msg, pagination
func JSONResponseWithPagination(c *gin.Context, code int, data interface{}, msg string, pagination Pagination) {
	if data == nil {
		data = errorData(code)
	}
	now := time.Now()
	var resp = gin.H{"code": code, "data": data, "msg": msg, "pagination": pagination, "server_time": now.Unix()}
	c.JSON(200, resp)
}
