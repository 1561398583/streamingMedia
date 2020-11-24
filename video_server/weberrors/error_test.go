package weberrors

import (
	"fmt"
	"testing"
)

type BZError struct {
	msg string
	code int
}

func (be *BZError) Error() string {
	return be.msg
}

func (be *BZError) Code() int {
	return be.code
}

type OtherError struct {
	msg string
}

func (e *OtherError) Error() string {
	return e.msg
}

func TestWebError(t *testing.T) {
	type StackAndMsg interface {
		Error() string
		GetStackInfo() string
	}
	err := warpError()
	if err != nil {
		switch e := Cause(err).(type) {
		case *BZError:	//如果是业务error，就只需打印出错误信息
			fmt.Println("<get a BZError>")
			fmt.Println(e)
		default:	//其他未知错误，就要打印出栈信息，方便分析
			if et, ok := err.(StackAndMsg); ok {
				//如果实现了StackAndMsg接口，*WebError就实现了这个接口，相当于判断err是否是*WebError类型
				fmt.Println("<get a StackAndMsg>")
				fmt.Println(et.Error())
				fmt.Println(et.GetStackInfo())
			}else {
				//否则，应该是其他第三方抛出的原始error
				fmt.Println("<get original error>")
				fmt.Println(err)
			}
		}
	}else {
		fmt.Println("no error")
	}

}

func warpError()  error{
	err := getError()
	if err != nil {
		return Wrap(err)
	}
	return nil
}

func getError() error {
	//err := New("get a error")
	//err := &BZError{msg: "user is not exist", code: 1}
	err := &OtherError{msg: "unknow error"}
	return err
}

