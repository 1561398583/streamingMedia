package errorOrPanic

type BuzErr struct {
	msg string
	code int
}

func (be *BuzErr) Error() string{
	return be.msg
}

func ErrorOrPanic(e error) error {
	if _, ok := e.(*BuzErr); ok {
		return e
	}else {
		panic(e)
	}
}

func New(msg string, code int)  *BuzErr{
	return &BuzErr{msg: msg, code: code}
}
