package weberrors

import (
	"runtime"
	"strconv"
)

type WebError struct {
	err error
	StackInfo string
	ErrorType int	//0:未知错误，需要记录错误位置到log；1.业务错误,请求对象没找到，只需返回错误信息给用户
}

type WE struct {
	msg string
}

func (we *WE) Error() string{
	return we.msg
}


//ErrorType
const (
	UNKNOW_ERROR = 0
	NOT_FOUND = 1
	PARA_ERROR = 2	//参数错误
)

func (we *WebError) Error() string {
	return we.err.Error()
}


func (we *WebError) GetStackInfo() string {
	return we.StackInfo
}


func New(msg string) *WebError {
	err := WE{msg: msg}
	stackInfo := getStack()
	return &WebError{err: &err, StackInfo: stackInfo}
}

//获取位置
func getLine() string {
	_, fileName, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file line")
	}
	return "[" + fileName + " : " + strconv.FormatInt(int64(line), 10) + "]"
}

//获取调用栈信息
func getStack() string {
	s := ""
	for i := 0; ;i++ {
		_, fileName, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		s = s + fileName + " " + strconv.FormatInt(int64(line), 10) + "\n"
	}
	return s
}

func Wrap(e error) error {
	//如果是WebError,直接返回;否则把error封装成WebError返回
	if we, ok := e.(*WebError); ok{
		return we
	}else {
		return &WebError{err: e, StackInfo: getStack()}
	}
}

//获取error类型
func Cause(e error) error {
	if we, ok := e.(*WebError); ok{
		//如果是*WebError，就把error提取出来
		return we.err
	}
	return e
}




