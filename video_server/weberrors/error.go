package weberrors

import (
	"runtime"
	"strconv"
)

type WebError struct {
	Msg string
	StackInfo string
	ErrorType int	//0:未知错误，需要记录错误位置到log；1.业务错误,请求对象没找到，只需返回错误信息给用户
}

//ErrorType
const (
	UNKNOW_ERROR = 0
	NOT_FOUND = 1
)

func (we *WebError) Error() string {
	return we.Msg + "\n" + we.StackInfo
}

func (we *WebError) GetType() int {
	return we.ErrorType
}

func (we *WebError) GetMsg() string {
	return we.Msg
}

func New(msg string, et int) *WebError {
	line := getLine()
	return &WebError{Msg: msg, StackInfo: line, ErrorType: et}
}

//获取位置
func getLine() string {
	_, fileName, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file line")
	}
	return "[" + fileName + " : " + strconv.FormatInt(int64(line), 10) + "]"
}

func Wrap(e error) error {
	//如果是WebError，则添加调用栈信息;否则把error封装成WebError返回
	if we, ok := e.(*WebError); ok{
		we.StackInfo = we.StackInfo + "\n" + getLine()
		return we
	}else {
		return New(e.Error(), UNKNOW_ERROR)
	}
}

func Type(e error) int {
	if we, ok := e.(*WebError); ok{
		return we.GetType()
	}else {
		return UNKNOW_ERROR
	}
}




