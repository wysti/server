package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

type RestHandler interface {
	Handler(http.ResponseWriter, *http.Request)
	ServeRest(params map[string]string) (interface{}, *RestError)
	Context() string
}

type RestHandlerBase struct {
	context string
	rh      RestHandler
}

func NewRestHandlerBase(context string, rh RestHandler) *RestHandlerBase {
	rhb := new(RestHandlerBase)
	rhb.context = context
	rhb.rh = rh
	return rhb
}

func (rhb *RestHandlerBase) Handler(rw http.ResponseWriter, r *http.Request) {
	content, re := rhb.rh.ServeRest(nil)
	if re != nil {
		io.WriteString(rw, ErrorMessages[GENERIC_SERVER])
		return
	}
	byteArr, err := json.Marshal(content)
	if err != nil {
		io.WriteString(rw, ErrorMessages[JSON_MARSHALLING])
		return
	}
	rw.Write(byteArr)
}

func (rhb *RestHandlerBase) Context() string {
	return rhb.context
}
