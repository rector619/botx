package notification

import (
	"github.com/SineChat/bot-ms/external"
	"github.com/SineChat/bot-ms/utility"
)

type RequestObj struct {
	Name         string
	Path         string
	Method       string
	SuccessCode  int
	RequestData  interface{}
	DecodeMethod string
	Logger       *utility.Logger
}

var (
	JsonDecodeMethod    string = "json"
	PhpSerializerMethod string = "phpserializer"
)

func (r *RequestObj) getNewSendRequestObject(data interface{}, headers map[string]string, urlprefix string) *external.SendRequestObject {
	return external.GetNewSendRequestObject(r.Logger, r.Name, r.Path, r.Method, urlprefix, r.DecodeMethod, headers, r.SuccessCode, data)
}
