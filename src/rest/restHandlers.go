package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type APIProvider func() HandlerSpec

type RestHandler func(map[string][]string) (interface{}, *RestError)

type HandlerSpec struct {
	Context   string
	ServeRest RestHandler
}

type HttpHandler func(http.ResponseWriter, *http.Request)

func MuxHandler(hs HandlerSpec) HttpHandler {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		e0 := r.ParseForm()
		if e0 != nil {
			io.WriteString(rw, ErrorMessages[REQUEST_PARSE])
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		content, e1 := hs.ServeRest(r.Form)
		if e1 != nil {
			io.WriteString(rw, ErrorMessages[GENERIC_SERVER])
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		byteArr, e2 := json.Marshal(content)
		if e2 != nil {
			io.WriteString(rw, ErrorMessages[JSON_MARSHALLING])
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(byteArr)
	}
}

/**
 * The methods of restAPI should all implement APIProvider, defined just above.
 * APIProvider returns
 * ** a HandlerSpec struct, which contains the http context on which the Handler will listen,
 * ** a RestHandler, which MuxHandler uses to respond to HttpRequests.
 * The idea is to isolate the API writer from HTTP specifics. (RestHandler takes in string parameters and returns an interface)
 */
func AcceptRequests(restAPI interface{}) {
	mux := http.NewServeMux()
	raValue := reflect.ValueOf(restAPI)
	for i := 0; i < raValue.NumMethod(); i++ {
		spec := raValue.Method(i).Call([]reflect.Value{})
		specIF := spec[0].Interface()
		hs := specIF.(HandlerSpec)
		handler := MuxHandler(hs)
		fmt.Println("handling", hs.Context)
		mux.HandleFunc(hs.Context, handler)
	}
	fmt.Println("Listening...")
	http.ListenAndServe(":9090", mux)
}
