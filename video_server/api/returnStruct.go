package api

type ResponseStruct struct {
	Status string	//"ok" or "error"
	Code int
	Data string	//json
}

const (
	OK = "ok"
	ERROR = "error"
)
