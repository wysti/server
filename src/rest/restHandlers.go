package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

type RestHandler interface {
	Handler(http.ResponseWriter, *http.Request)
	ServeRest(params map[string][]string) (interface{}, *RestError)
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
	e0 := r.ParseForm()
	if e0 != nil {
		io.WriteString(rw, ErrorMessages[REQUEST_PARSE])
		return
	}
	content, e1 := rhb.rh.ServeRest(r.Form)
	if e1 != nil {
		io.WriteString(rw, ErrorMessages[GENERIC_SERVER])
		return
	}
	byteArr, e2 := json.Marshal(content)
	if e2 != nil {
		io.WriteString(rw, ErrorMessages[JSON_MARSHALLING])
		return
	}
	rw.Write(byteArr)
}

func (rhb *RestHandlerBase) Context() string {
	return rhb.context
}
